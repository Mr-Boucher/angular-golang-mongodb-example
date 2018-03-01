package main

import (

	"crypto/tls"
	"net"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type TestData struct {
	Id         bson.ObjectId   `bson:"_id,omitempty" json:"-"`
	Value      string   `bson:"value" json:"value"`

}

//Kick it all off
func main() {
	loadData();
	router := mux.NewRouter()
	router.HandleFunc("/", GetData).Methods("GET")
	router.HandleFunc("/loaddata", GetData).Methods("GET", "OPTIONS")//.Headers("Content-Type", "application/json")
	log.Fatal(http.ListenAndServe(":8000", router))
}

//Get the data from mongo
func GetData(writer http.ResponseWriter, request *http.Request) {

	writer.Header().Add("Access-Control-Allow-Origin", "*") //allow access from anywhere
	writer.Header().Add("Access-Control-Allow-Methods", "GET")
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	writer.Header().Add("Content-Type", "application/json")

	fmt.Println( "Path:", request.URL.Path )
	data := loadData()
	json.NewEncoder(writer).Encode(data)
}

//Load data from mongo
func loadData( ) []TestData {

	//todo pass in
	databaseName := "dev"
	collectionName := "test"

	//set connection to mongo
	//todo pass in
	dialInfo := &mgo.DialInfo{
		Addrs: []string{"cluster0-shard-00-00-iaz9w.mongodb.net:27017"},
		Database: "admin",
		Username: "dev",
		Password: "dev",
	}

	fmt.Println( "Opening connection to", dialInfo.Addrs, "as", dialInfo.Username, "from the", dialInfo.Database, "DB." )
	fmt.Println( "Start loading data" )

	//todo figure out what this is
	tlsConfig := &tls.Config{}

	//call the mongo server
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}

	//Create the session
	fmt.Println( "Creating session" )
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil { panic(err) } //fail in null

	//Get the mongo db from the session
	fmt.Println( "Opening DB", databaseName, "using session", session )
	database := session.DB(databaseName)
	if database == nil { panic(database) }

	//Get the collection
	collection := database.C( collectionName )
	if collection == nil {	panic(collection) }

	//Load data
	query := collection.Find( bson.M{} )
	count, err := query.Count()
	if err != nil {panic( err)	}

	//
	fmt.Println( "Count:", count)
	results := []TestData{}
	query.All( &results )
	for id, result := range results {
		fmt.Println( id, "id:", result.Id )
		fmt.Println( "value:", result.Value )
	}

	fmt.Println( "Finished loading data" )

	return results
}

//org test
func test() {
	fmt.Println( "Start" )
	tlsConfig := &tls.Config{}

	//
	dialInfo := &mgo.DialInfo{
		Addrs: []string{"cluster0-shard-00-00-iaz9w.mongodb.net:27017"},
		Database: "admin",
		Username: "dev",
		Password: "dev",
	}

	//
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		panic(err)
	}

	fmt.Println( "session:", session )
	database := session.DB("dev")

	names, err := database.CollectionNames()
	if err != nil {
		panic(err)
	}

	for _, name := range names {
		// index is the index where we are
		// element is the element from someSlice for where we are
		fmt.Println( "Getting Collec  tion Name:", name )
		collection := database.C( name )
		fmt.Println( "Got:", collection.Name )
		query := collection.Find( bson.M{} )
		count, err := query.Count()
		if err != nil {
			panic( err)
		}

		fmt.Println( "Count:", count)
		results := []TestData{}
		query.All( &results )
		for id, result := range results {
			fmt.Println( id, "id:", result.Id )
			fmt.Println( "value:", result.Value )
		}
	}

	fmt.Println( "End" )
}

