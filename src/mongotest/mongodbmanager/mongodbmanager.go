package mongodbmanager

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"crypto/tls"
	"net"
	"time"
)

type ActionArgument interface{} //arguments for different actions
type ActionResults interface{}  //results from different actions

//
type MongoDBManager struct {
	configuration MongoDBConfiguration
}

//Constructor
func NewMongoDBManager(configuration MongoDBConfiguration) *MongoDBManager {
	m := MongoDBManager{}
	fmt.Println( "MongoDBManager::NewMongoDBManager", m )
	return &m
}

//
func (db *MongoDBManager) InitContext(contextHolder ContextHolder) {

	//
	context := contextHolder.GetMongoDBContext()
	fmt.Println( "MongoDBManager::InitContext MongoDBContext:", context);

	databaseConnectionInfo := context.GetConfiguration()
	fmt.Println( "MongoDBManager::InitContext databaseConnectionInfo:", databaseConnectionInfo );
	if databaseConnectionInfo == nil {
		panic( "Missing databaseConnectionInfo is nil" )
	}

	//set connection to mongo
	dialInfo := &mgo.DialInfo{
		Addrs:    databaseConnectionInfo.Cluster,
		Database: databaseConnectionInfo.UserDatabase,
		Username: databaseConnectionInfo.Username,
		Password: databaseConnectionInfo.Password,
	}

	fmt.Println("Opening connection to", dialInfo.Addrs, "as", dialInfo.Username, "from the", dialInfo.Database, "DB.")

	//call the mongo server
	tlsConfig := &tls.Config{} //todo figure out what this is
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}

	//Create the session
	fmt.Println("Creating session:")
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Created session:", session)
	context.SetSession( session )
}

//Clean up any resources created by the InitContext
func (db *MongoDBManager) CleanupContext(contextHolder ContextHolder) {
	session := contextHolder.GetMongoDBContext().GetSession()
	fmt.Println("Closing session:", session)
	session.Close()
}

//Method handling framework calls to mongoDB this method will create and destroy all resources needed
//to work with mongoDB it will perform the action function and return the results
func (db *MongoDBManager) Execute(contextHolder ContextHolder, action func(context interface{}, arguments interface{}) interface{}, arguments interface{}) interface{} {
	startTime := float64(time.Now().UnixNano() / int64(time.Millisecond))
	fmt.Println("MongoDBManager::Context:", contextHolder )

	//
	databaseConnectionInfo := contextHolder.GetMongoDBContext().GetConfiguration()

	//Get the mongo db from the session
	session := contextHolder.GetMongoDBContext().GetSession()
	fmt.Println("Opening DB", databaseConnectionInfo.DatabaseName, "using session", session)
	database := session.DB(databaseConnectionInfo.DatabaseName)
	if database == nil {
		panic(database)
	}

	//Get the collection
	collection := database.C(databaseConnectionInfo.CollectionName)
	if collection == nil {
		panic(collection)
	}

	//Open the monitoring collection, create is does not exist
	//monitorCollection := database.C("monitoring")
	//if collection == nil {
	//	fmt.Println( "Creating monitoring collection" )
	//}

	//Wrap the collection to decorate it with other features
	//wrapper := NewCollectionWrapper( collection, nil, nil )

	//execute the acton function
	contextHolder.GetMongoDBContext().SetCollection( collection )
	result := action(contextHolder, arguments)

	endTime := float64(time.Now().UnixNano() / int64(time.Millisecond))
	duration := endTime - startTime
	//fmt.Println( "Query", query , "took", duration )
	fmt.Println( "Query" , "took", duration )

	return result;
}

