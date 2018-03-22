package dataeditor

import (
	"fmt"
	"reflect"

	"gopkg.in/mgo.v2/bson"
	"github.com/rs/xid"

	"encoding/json"

	"../httpmanager"
	"../mongodbmanager"
)

const(
	baseUrl = "/data"
)

//
type Context interface {
	GetParameters() map[string]string
	GetSliceParameters() map[string][]string
	GetCollection() mongodbmanager.CollectionWrapper
}

//
type DataEditor interface {
	GetId() int
	SetId( int )
	GetHttpRouterHandlers() []httpmanager.HttpRouterHandler
	//Marshal( theData interface {} ) ([]byte, error)
	Unmarshal( []byte ) (interface {}, error)
}

//
type dataEditorObject struct {
	id int
}

//Test data format
type TestData struct {
	ObjectId bson.ObjectId `bson:"_id,omitempty" json:"-"` //Setup mapping for data from bson(used for Mongo) and json(Used by REST API response) note the "-" it means that json does not have this
	Value    string        `bson:"value" json:"value"`     //Setup mapping for data from bson(used for Mongo) and json(Used by REST API response)
	Id       string        `bson:"id" json:"id"`           //Setup mapping for data from bson(used for Mongo) and json(Used by REST API response)
}

//
func NewEditor( ) DataEditor {
	config := dataEditorObject{}
	return &config
}

//GetHttpRouterHandlers() meets the interface
func (d *dataEditorObject) GetHttpRouterHandlers() []httpmanager.HttpRouterHandler {

	//Create the call back methods for /Data
	dataMethodFunctions := []httpmanager.HttpMethodFunction{}
	dataMethodFunctions = append( dataMethodFunctions, httpmanager.NewHttpMethodFunction( "GET", d.search ) )
	dataMethodFunctions = append( dataMethodFunctions, httpmanager.NewHttpMethodFunction( "POST", d.create ) )
	dataMethodFunctions = append( dataMethodFunctions, httpmanager.NewHttpMethodFunction( "DELETE", d.deleteAll ) )
	dataMethodHandler := httpmanager.NewHttpRouteHandler( d.id, baseUrl, dataMethodFunctions )

	//Search
	dataSearchMethodFunctions := []httpmanager.HttpMethodFunction{}
	dataSearchMethodFunctions = append( dataSearchMethodFunctions, httpmanager.NewHttpMethodFunction( "GET", d.search ) )
	dataSearchMethodHandler := httpmanager.NewHttpRouteHandler( d.id, baseUrl + "?search=[a-zA-Z0-9]+}", dataSearchMethodFunctions )

	//add the backs method for /Data/id
	dataIdMethodFunctions := []httpmanager.HttpMethodFunction{}
	dataIdMethodFunctions = append( dataIdMethodFunctions, httpmanager.NewHttpMethodFunction( "PUT", d.update ) )
	dataIdMethodFunctions = append( dataIdMethodFunctions, httpmanager.NewHttpMethodFunction( "DELETE", d.deleteById ) )
	dataIdMethodHandler := httpmanager.NewHttpRouteHandler( d.id, baseUrl + "/{id:[a-z0-9]+}", dataIdMethodFunctions )

	//
	routers := []httpmanager.HttpRouterHandler{}
	routers = append( routers, dataMethodHandler )
	routers = append( routers, dataSearchMethodHandler )
	routers = append( routers, dataIdMethodHandler )

	return routers
}

//
func (d *dataEditorObject) Unmarshal( payload []byte ) (interface {}, error) {

	theData := TestData{}
	err := json.Unmarshal(payload, &theData)

	fmt.Println("Data Editor:", theData )

	return theData, err
}

func (d *dataEditorObject) GetId() int {
	return d.id
}

func (d *dataEditorObject) SetId( id int ) {
	d.id = id
}

//Load data from mongo returned as a []TestData
func (d *dataEditorObject) search( appcontext interface{}, arguments interface{} ) interface{} {

	fmt.Println( "DataEditor::Search", arguments )
	var results []TestData

	context := appcontext.(Context)

	//Get the search criteria
	searchCriteria := context.GetSliceParameters()["search"]
	fmt.Println( "DataEditor::Search searchCriteria", searchCriteria )

	//Load data
	collection := context.GetCollection()
	var criteria bson.M
	if len(searchCriteria) > 0 {
		if len(searchCriteria[0]) > 0 {
			fmt.Println( "DataEditor::Search searchCriteria", searchCriteria[0] )
			regex := bson.RegEx{}
			regex.Pattern = "^" + searchCriteria[0]
			regex.Options = "i"
			criteria = bson.M{"value": regex }
		}
	}
	query := collection.Find(criteria)
	query = query.Sort("value") //sort the data by its value

	query.All(&results) //execute the query

	//Display the data returned for debugging
	for index, result := range results {
		fmt.Println(index, "id:", result.Id, "value:", result.Value)
	}

	fmt.Println("Finished Search data")

	return results
}

////remove data from db base
func (d *dataEditorObject) create(appcontext interface{}, arguments interface{} ) interface{} {
	fmt.Println( "DataEditor::create arguments", arguments )
	context := appcontext.(Context)

	//Validation handling
	if arguments == nil {
		panic("Missing argument of type TestData")
	}

	fmt.Println("arguments:", arguments)

	//Convert the empty interface to a string that contains the id
	newData, ok := arguments.(TestData) //same as casting in java
	if !ok {
		errorMessage := fmt.Sprint("Argument should be of type TestData. It was ", reflect.TypeOf(arguments))
		panic(errorMessage)
	}

	//make sure we create the id before storing it
	newData.Id = xid.New().String()

	//Insert the TestData
	fmt.Println("create:", newData)
	collection := context.GetCollection()
	err := collection.Insert(newData)
	if err != nil {
		panic(err)
	}

	//creating new one the first element is the only one that needs a new id
	fmt.Println("create:", "finished", newData)

	return newData
}

//remove data from db base
func (d *dataEditorObject) update(context interface{}, arguments interface{} ) interface{} {
	fmt.Println( "DataEditor::update arguments", arguments )

	//id, ok := arguments.(string)
	//if !ok {
	//	panic( "Argument should be of type string" )
	//}
	//
	//collection := context.(*mgo.Collection)
	//collection.Remove( bson.M{"id": id} )
	fmt.Println("update:", "finished")
	return nil
}

//remove data from db base
func (d *dataEditorObject) deleteById(appcontext interface{}, arguments interface{} ) interface{} {
	context := appcontext.(Context)
	id := context.GetParameters()["id"] //get the id of the object to delete
	fmt.Println( "DataEditor::deleteById", context.GetParameters(), "started" )

	collection := context.GetCollection()
	fmt.Println( "DataEditor::deleteById collection", collection)
	err := collection.Remove( bson.M{"id": id} )
	if err != nil {
		panic( err )
	}
	fmt.Println( "DataEditor::deleteById", id, "finished" )
	return nil
}

//remove data from db base
func (d *dataEditorObject) deleteAll(appcontext interface{}, arguments interface{} ) interface{} {

	fmt.Println( "DataEditor::deleteAll arguments", arguments )
	//context := appcontext.(Context)
	//id := context.GetParameters()["id"] //get the id of the object to delete
	//
	//fmt.Println( "arguments:", arguments )
	//
	////Validation handling
	//if arguments == nil {
	//	panic( "Missing argument of type string" )
	//}
	//
	////Convert the empty interface to a string that contains the id
	//id, ok := arguments.(string)
	//if !ok {
	//	panic( "Argument should be of type string" )
	//}
	//
	//fmt.Println("deleteById:", id, "started")
	//collection := context.GetCollection()
	//err := collection.Remove( bson.M{"id": id} )
	//if err != nil {
	//	panic( err )
	//}
	//fmt.Println("deleteById:", id, "finished")
	return nil
}
