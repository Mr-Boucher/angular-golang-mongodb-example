package configuration

import (
	"fmt"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"../mongodbmanager"

	"../httpmanager"
)

const(
	baseUrl = "/configuration"
)

//
type Context interface {
	GetParameters() map[string]string
	GetSliceParameters() map[string][]string
	GetCollection() mongodbmanager.CollectionWrapper
}

//Test data format
type ConfigurationData struct {
	ObjectId bson.ObjectId `bson:"_id,omitempty" json:"-"` //Setup mapping for data from bson(used for Mongo) and json(Used by REST API response) note the "-" it means that json does not have this
	Value    string        `bson:"value" json:"value"`     //Setup mapping for data from bson(used for Mongo) and json(Used by REST API response)
	Id       string        `bson:"id" json:"id"`           //Setup mapping for data from bson(used for Mongo) and json(Used by REST API response)
}

//
type Configuration interface {
	GetId() int
	SetId( int )
	GetHttpRouterHandlers() []httpmanager.HttpRouterHandler
	//Marshal( theData interface {} ) ([]byte, error)
	Unmarshal( []byte ) (interface {}, error)
}

//
type ConfigurationObject struct {
	id int
}

//
func NewConfiguration( ) Configuration {
	config := ConfigurationObject{}
	return &config
}

func (d *ConfigurationObject) GetId() int {
	return d.id
}

func (d *ConfigurationObject) SetId( id int ) {
	d.id = id
}

//GetHttpRouterHandlers() meets the interface
func (d *ConfigurationObject) GetHttpRouterHandlers() []httpmanager.HttpRouterHandler {

	//Create the call back methods for /Data
	dataMethodFunctions := []httpmanager.HttpMethodFunction{}
	dataMethodFunctions = append( dataMethodFunctions, httpmanager.NewHttpMethodFunction( "GET", d.load ) )
	dataMethodHandler := httpmanager.NewHttpRouteHandler( d.id, baseUrl, dataMethodFunctions )


	//add the backs method for /Data/id
	dataIdMethodFunctions := []httpmanager.HttpMethodFunction{}
	dataIdMethodFunctions = append( dataIdMethodFunctions, httpmanager.NewHttpMethodFunction( "PUT", d.update ) )
	dataIdMethodHandler := httpmanager.NewHttpRouteHandler( d.id, baseUrl + "/{id:[a-z0-9]+}", dataIdMethodFunctions )

	//
	routers := []httpmanager.HttpRouterHandler{}
	routers = append( routers, dataMethodHandler )
	routers = append( routers, dataIdMethodHandler )

	return routers
}

//
func (d *ConfigurationObject) Unmarshal( payload []byte ) (interface {}, error) {

	theData := ConfigurationData{}
	err := json.Unmarshal(payload, &theData)

	fmt.Println("ConfigurationData:", theData )

	return theData, err
}

//
func (d *ConfigurationObject) load( appcontext interface{}, arguments interface{} ) (interface{}, error) {
	fmt.Println( "ConfigurationData::load arguments", arguments )
	var results []ConfigurationData
	var err error


	//todo fill configuration data
	results = append(results, ConfigurationData{"ObjectId", "Value", "Id"})

	//context := appcontext.(Context)
	//collection := context.GetCollection()
	//
	////Load data
	//query := collection.Find(bson.M{})
	//query = query.Sort("value") //sort the data by its value
	//
	//query.All(&results) //execute the query
	//
	////Display the data returned for debugging
	//for index, result := range results {
	//	fmt.Println(index, "id:", result.Id, "value:", result.Value)
	//}

	fmt.Println("Finished loading data")

	return results, err
}

////remove data from db base
func (d *ConfigurationObject) update(context interface{}, arguments interface{} ) (interface{}, error) {
	fmt.Println( "ConfigurationData::update arguments", arguments )
	var err error

	//id, ok := arguments.(string)
	//if !ok {
	//	panic( "Argument should be of type string" )
	//}
	//
	//collection := context.(*mgo.Collection)
	//collection.Remove( bson.M{"id": id} )
	fmt.Println("update:", "finished")
	return nil, err
}


