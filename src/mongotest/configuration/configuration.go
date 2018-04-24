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
	var err error

	context := appcontext.(mongodbmanager.ContextHolder)
	info := context.GetMongoDBContext()
	config := info.GetConfiguration()
	fmt.Println(config)

	fmt.Println("Finished loading data")

	return config, err
}

////remove data from db base
func (d *ConfigurationObject) update(context interface{}, arguments interface{} ) (interface{}, error) {
	fmt.Println( "ConfigurationData::update arguments", arguments )
	var err error

	contextHolder := context.(mongodbmanager.ContextHolder)
	mongo := contextHolder.GetMongoDBContext()
	config := mongo.GetConfiguration()

	fmt.Println(config)

	fmt.Println("update:", "finished")
	return nil, err
}


