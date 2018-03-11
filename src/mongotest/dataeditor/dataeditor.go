package dataeditor

import (
	"fmt"
	"reflect"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/rs/xid"

	"../httpmanager"
	"encoding/json"
)

const(
	baseUrl = "/data"
)

//
type DataEditor struct {
	id int
}

//Test data format
type TestData struct {
	ObjectId bson.ObjectId `bson:"_id,omitempty" json:"-"` //Setup mapping for data from bson(used for Mongo) and json(Used by REST API response) note the "-" it means that json does not have this
	Value    string        `bson:"value" json:"value"`     //Setup mapping for data from bson(used for Mongo) and json(Used by REST API response)
	Id       string        `bson:"id" json:"id"`           //Setup mapping for data from bson(used for Mongo) and json(Used by REST API response)
}

//
func (d *DataEditor) GetHttpRouterHandlers() []httpmanager.HttpRouterHandler {
	routers := []httpmanager.HttpRouterHandler{} //empty array for routers

	//add /data route for GET and POST
	funcs := []httpmanager.HttpMethodFunction{{"GET", d.load}, {"POST", d.create} }
	routers = append( routers, httpmanager.HttpRouterHandler{ d.id, baseUrl, funcs } )

	funcs = []httpmanager.HttpMethodFunction{{"PUT", d.update}, {"DELETE", d.deleteById} }
	routers = append( routers, httpmanager.HttpRouterHandler{ d.id, baseUrl + "/{id:[a-z0-9]+}", funcs } )

	return routers
}

//
func (d *DataEditor) Unmarshal( payload []byte ) (interface {}, error) {

	theData := TestData{}
	err := json.Unmarshal(payload, &theData)

	fmt.Println("Data Editor:", theData )

	return theData, err
}

func (d *DataEditor) GetId() int {
	return d.id
}

func (d *DataEditor) SetId( id int ) {
	d.id = id
}

//Load data from mongo returned as a []TestData
func (d *DataEditor) load( context interface{}, arguments interface{} ) interface{} {

	fmt.Println("dataeditor.load")
	var results []TestData

	//Load data
	collection := context.(*mgo.Collection)
	query := collection.Find(bson.M{})
	query = query.Sort("value") //sort the data by its value

	query.All(&results) //execute the query

	//Display the data returned for debugging
	for index, result := range results {
		fmt.Println(index, "id:", result.Id, "value:", result.Value)
	}

	fmt.Println("Finished loading data")

	return results
}

////remove data from db base
func (d *DataEditor) create(context interface{}, arguments interface{} ) interface{} {

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

	//Insert the TestData
	fmt.Println("create:", "started")
	collection := context.(*mgo.Collection)
	err := collection.Insert(newData)
	if err != nil {
		panic(err)
	}

	//creating new one the first element is the only one that needs a new id

	newData.Id = xid.New().String()

	fmt.Println("create:", "finished", newData)

	return newData
}

//
////remove data from db base
func (d *DataEditor) update(context interface{}, arguments interface{} ) interface{} {

	fmt.Println("update:", "started")
	id, ok := arguments.(string)
	if !ok {
		panic( "Argument should be of type string" )
	}

	collection := context.(*mgo.Collection)
	collection.Remove( bson.M{"id": id} )
	fmt.Println("update:", "finished")
	return nil
}
//
////remove data from db base
func (d *DataEditor) deleteById(context interface{}, arguments interface{} ) interface{} {

	//Validation handling
	if arguments == nil {
		panic( "Missing argument of type string" )
	}

	//Convert the empty interface to a string that contains the id
	id, ok := arguments.(string)
	if !ok {
		panic( "Argument should be of type string" )
	}

	fmt.Println("deleteById:", id, "started")
	collection := context.(*mgo.Collection)
	err := collection.Remove( bson.M{"id": id} )
	if err != nil {
		panic( err )
	}
	fmt.Println("deleteById:", id, "finished")
	return nil
}
