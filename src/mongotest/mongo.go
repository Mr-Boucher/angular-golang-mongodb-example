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
	"strings"
	"os"
)

//Test data format
type TestData struct {
	ObjectId bson.ObjectId `bson:"_id,omitempty" json:"-"` //Setup mapping for data from bson(used for Mongo) and json(Used by REST API response) note the "-" it means that json does not have this
	Value    string        `bson:"value" json:"value"`     //Setup mapping for data from bson(used for Mongo) and json(Used by REST API response)
	Id       string        `bson:"id" json:"id"`           //Setup mapping for data from bson(used for Mongo) and json(Used by REST API response)
}

type MongoDB struct {
	databaseName   string
	collectionName string
	cluster        []string
	userDatabase   string
	username       string
	password       string
}

//Should be a constant but can't because of language restriction that const can't have arrays
var mongoDB = MongoDB{"dev", "test", []string{"cluster0-shard-00-00-iaz9w.mongodb.net:27017"}, "admin", "dev", "dev"}

//Kick it all off
func main() {
	//Log start up arguments
	fmt.Println(strings.Join(os.Args, " "))

	//Find the start up port
	port := "8000"; //default port
	for _, arg := range os.Args[1:] {
		arg := strings.Split(arg, "=")
		if arg[0] == "port" {
			port = arg[1]
		}
	}
	fmt.Println("Starting up on", port)

	//Setup route for incoming data requests
	router := mux.NewRouter()
	router.HandleFunc("/data", options).Methods("OPTIONS")             //Setup data the REST API and call options
	router.HandleFunc("/data/{id:[0-9]+}", options).Methods("OPTIONS") //Setup data the REST API and options

	router.HandleFunc("/data", getData).Methods("GET")                   //Setup data as the REST API and call GetData for get requests
	router.HandleFunc("/data", createData).Methods("POST")               //Setup data the REST API and call CreateData for delete requests
	router.HandleFunc("/data/{id:[0-9]+}", updateData).Methods("PUT")    //Setup data the REST API and call UpdateData for delete requests
	router.HandleFunc("/data/{id:[0-9]+}", deleteData).Methods("DELETE") //Setup data the REST API and call DeleteData for delete requests

	//Start listening for requests - thread waits forever at this port
	log.Fatal(http.ListenAndServe(":"+port, router))
}

////////////////////////////////////////REST API FUNCTIONS/////////////////////////////////////////////////////////////////////////
//Set headers to tell the client what is supported for this REST API
func setHeaders(writer http.ResponseWriter) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")                                            //Allow access from anywhere
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Request-Origin") //Allows setting of the Content-Type by the client
	writer.Header().Add("Access-Control-Allow-Methods", "HEAD, GET, POST, PUT, DELETE, OPTIONS")       //REST API supports GET, POST, PUT, DELETE
	writer.Header().Add("Accept", "application/json")                                                  //Only json is accepted
	writer.Header().Add("Content-Type", "application/json")
}

//Called before other REST requests to make sure all the headers are correctly set
//As well as set up security headers
func options(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("options")
	setHeaders(writer)
}

//Get the data from mongo
func getData(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("getData:", request.URL.Path)
	setHeaders(writer) //Set response headers

	data := execute(mongoDB, load, nil)

	json.NewEncoder(writer).Encode(data) //stream the encoded data on the writer
}

//
func createData(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("createData:", request.URL.Path)
	setHeaders(writer) //Set response headers
	execute(mongoDB, create, nil)
}

//
func updateData(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("updateData:", request.URL.Path)
	setHeaders(writer) //Set response headers
	execute(mongoDB, update, nil)
}

//
func deleteData(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request) //retrieve the query params from the url
	id := params["id"] //get the id of the object to delete
	fmt.Printf("deleteData(%s):%s\n", id, request.URL.Path)
	setHeaders(writer) //Set response headers

	//perform the deleteDataById
	//args: actionArgument = id
	execute(mongoDB, deleteById, id)
}

////////////////////////////////////////ACTION FUNCTIONS/////////////////////////////////////////////////////////////////////////
//Load data from mongo returned as a []TestData
func load(collection *mgo.Collection, argument actionArgument) actionResults {

	//Load data
	query := collection.Find(bson.M{})
	query = query.Sort("value") //sort the data by its value
	var results []TestData
	query.All(&results) //execute the query

	//Display the data returned for debugging
	for index, result := range results {
		fmt.Println(index, "id:", result.Id, "value:", result.Value)
	}

	fmt.Println("Finished loading data")

	return results
}

//remove data from db base
func create(collection *mgo.Collection, argument actionArgument) actionResults {

	//Validation handling
	if argument == nil {
		panic( "Missing argument of type TestData" )
	}

	//Convert the empty interface to a string that contains the id
	newData, ok := argument.(TestData) //same as casting in java
	if !ok {
		panic( "Argument should be of type TestData" )
	}

	//Insert the TestData
	fmt.Println("create:", "started")
	err := collection.Insert( newData )
	if err != nil {
		panic( err )
	}
	fmt.Println("create:", "finished")
	return nil
}

//remove data from db base
func update(collection *mgo.Collection, argument actionArgument) actionResults {

	fmt.Println("update:", "started")
	//collection.Remove( bson.M{"id": id} )
	fmt.Println("update:", "finished")
	return nil
}

//remove data from db base
func deleteById(collection *mgo.Collection, argument actionArgument) actionResults {

	//Validation handling
	if argument == nil {
		panic( "Missing argument of type string" )
	}

	//Convert the empty interface to a string that contains the id
	id, ok := argument.(string)
	if !ok {
		panic( "Argument should be of type string" )
	}

	fmt.Println("deleteById:", id, "started")
	err := collection.Remove( bson.M{"id": id} )
	if err != nil {
		panic( err )
	}
	fmt.Println("deleteById:", id, "finished")
	return nil
}

////////////////////////////////////////FRAMEWORK/////////////////////////////////////////////////////////////////////////

type actionArgument interface {} //arguments for different actions
type actionResults interface {}  //results from different actions

//Method handling framework calls to mongoDB this method will create and destroy all resources needed
//to work with mongoDB it will perform the action function and return the results
func execute(databaseConnectionInfo MongoDB, action func(collection *mgo.Collection, arguments actionArgument) actionResults, arguments actionArgument) actionResults {

	//set connection to mongo
	dialInfo := &mgo.DialInfo{
		Addrs:    databaseConnectionInfo.cluster,
		Database: databaseConnectionInfo.userDatabase,
		Username: databaseConnectionInfo.username,
		Password: databaseConnectionInfo.password,
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
	fmt.Println("Opening DB", databaseConnectionInfo.databaseName, "using session", session)
	database := session.DB(databaseConnectionInfo.databaseName)
	if database == nil {
		panic(database)
	}

	//Get the collection
	collection := database.C(databaseConnectionInfo.collectionName)
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
