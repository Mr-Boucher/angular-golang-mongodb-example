package dataeditor

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"gopkg.in/mgo.v2"
	"../httphandler"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/rs/xid"
	"io/ioutil"
	"net/http"
)

const(
	baseUrl = "/data"
)

type DataEditor struct {

}

//Test data format
type TestData struct {
	ObjectId bson.ObjectId `bson:"_id,omitempty" json:"-"` //Setup mapping for data from bson(used for Mongo) and json(Used by REST API response) note the "-" it means that json does not have this
	Value    string        `bson:"value" json:"value"`     //Setup mapping for data from bson(used for Mongo) and json(Used by REST API response)
	Id       string        `bson:"id" json:"id"`           //Setup mapping for data from bson(used for Mongo) and json(Used by REST API response)
}

func (D DataEditor) New() *DataEditor {

}

//
func initializeHTTPSupport( handler httphandler.HttpHandler ) {
	//add /data route for GET and POST
	handler.Add( httphandler.HttpRouterHandler{baseUrl,
		{ httphandler.HttpMethodFunction{"GET", getData},
			httphandler.HttpMethodFunction{"POST", createData}},
		options} )

	//add /data/{id:[a-z0-9]+} for PUT and DELETE
	handler.Add( httphandler.HttpRouterHandler{baseUrl + "/{id:[a-z0-9]+}",
		{ httphandler.HttpMethodFunction{"PUT", updateData},
			httphandler.HttpMethodFunction{"DELETE", deleteById}},
		options} )
}

////////////////////////////////////////REST API FUNCTIONS/////////////////////////////////////////////////////////////////////////
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

	data := execute(configuration.MongoDB, load, nil)

	json.NewEncoder(writer).Encode(data) //stream the encoded data on the writer
}

//
func createData(writer http.ResponseWriter, request *http.Request) {
	// Read body
	body, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close() //make sure we clean up the steam
	if err != nil {
		http.Error(writer, err.Error(), 500)
		fmt.Println(err)
		return
	}

	var newData TestData
	fmt.Println( body )
	err = json.Unmarshal(body, &newData)
	if err != nil {
		http.Error(writer, err.Error(), 500)
		fmt.Println(err)
		return
	}


	fmt.Println("createData:", request.URL.Path)
	setHeaders(writer) //Set response headers

	newData.Id = xid.New().String()
	result := execute(configuration.MongoDB, create, newData )
	byteData, err := json.Marshal( result )
	writer.Write( byteData )
}

//
func updateData(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("updateData:", request.URL.Path)
	setHeaders(writer) //Set response headers
	execute(configuration.MongoDB, update, nil)
}

//
func deleteData(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request) //retrieve the query params from the url
	id := params["id"] //get the id of the object to delete
	fmt.Printf("deleteData(%s):%s\n", id, request.URL.Path)
	setHeaders(writer) //Set response headers

	//perform the deleteDataById
	//args: actionArgument = id
	execute(configuration.MongoDB, deleteById, id)
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
	fmt.Println("create:", "finished", newData)
	return newData
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
