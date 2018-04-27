package configuration

import (
	"fmt"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"../mongodbmanager"

	"../httpmanager"
	"reflect"
	"strconv"
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

type ConfigurationDataUpdate struct {
	ObjectId bson.ObjectId `bson:"_id,omitempty" json:"-"` //Setup mapping for data from bson(used for Mongo) and json(Used by REST API response) note the "-" it means that json does not have this
	Value    ConfigurationData        `bson:"value" json:"properties"`     //Setup mapping for data from bson(used for Mongo) and json(Used by REST API response)
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
	dataIdMethodHandler := httpmanager.NewHttpRouteHandler( d.id, baseUrl + "/{id:[A-Za-z0-9]+}", dataIdMethodFunctions )

	//
	routers := []httpmanager.HttpRouterHandler{}
	routers = append( routers, dataMethodHandler )
	routers = append( routers, dataIdMethodHandler )

	return routers
}

//
func (d *ConfigurationObject) Unmarshal( payload []byte ) (interface {}, error) {

	theData := ConfigurationDataUpdate{}
	err := json.Unmarshal(payload, &theData)

	fmt.Println("ConfigurationDataUpdate:", theData )

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
	fmt.Println( "ConfigurationDataUpdate::update arguments", arguments )
	var err error

	contextHolder := context.(mongodbmanager.ContextHolder)
	mongo := contextHolder.GetMongoDBContext()
	config := mongo.GetConfiguration()

	updateObject, ok := arguments.(ConfigurationDataUpdate)

	if !ok {
		errorMessage := fmt.Sprint("Argument should be of type ConfigurationDataUpdate. It was ", reflect.TypeOf(arguments))
		panic(errorMessage)
	}

	fmt.Println(updateObject)
	//missing database cluster right now
	if (updateObject.Id == "DatabaseName") {
		config.DatabaseName = updateObject.Value.Value
	} else if (updateObject.Id == "CollectionName") {
		config.CollectionName = updateObject.Value.Value
	} else if (updateObject.Id == "UserDatabase") {
		config.UserDatabase = updateObject.Value.Value
	} else if (updateObject.Id == "Username") {
		config.Username = updateObject.Value.Value
	} else if (updateObject.Id == "Password") {
		config.Password = updateObject.Value.Value
	} else if (updateObject.Id == "ConnectionTimeOut") {
		connectionTimeOut, err := strconv.Atoi(updateObject.Value.Value)
		if err != nil {
			config.ConnectionTimeOut = connectionTimeOut
		} else {
			errorMessage := fmt.Sprint(err)
			panic(errorMessage)
		}
	}

	fmt.Println(config)

	fmt.Println("update:", "finished")
	return nil, err
}


