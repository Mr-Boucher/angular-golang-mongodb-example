package mongodbmanager

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"crypto/tls"
	"net"
)

type ActionArgument interface{} //arguments for different actions
type ActionResults interface{}  //results from different actions

//
type MongoDBManager struct {
	Configuration MongoDBConfiguration
}

//
type MongoDBConfiguration struct {
	DatabaseName   string
	CollectionName string
	Cluster        []string
	UserDatabase   string
	Username       string
	Password       string
}

//Method handling framework calls to mongoDB this method will create and destroy all resources needed
//to work with mongoDB it will perform the action function and return the results
func (db *MongoDBManager) Execute( action func(collection *mgo.Collection, arguments ActionArgument) ActionResults, arguments ActionArgument) ActionResults {

	databaseConnectionInfo := MongoDBConfiguration{}

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

	//If we have a valid session make sure it is closed once the method exits otherwise it will be a session leak
	defer cleanUp( session )

	//Get the mongo db from the session
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

	//execute the acton function
	result := action(collection, arguments)

	return result;
}

//Make sure the resources are cleaned up
func cleanUp(session *mgo.Session) {
	fmt.Println( "Closing session:", session)
	session.Close()
}
