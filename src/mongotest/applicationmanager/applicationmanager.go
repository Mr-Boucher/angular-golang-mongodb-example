package applicationmanager

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../httpmanager"
	"../mongodbmanager"

)

//
type ManagerContext struct {
	HttpConnection httpmanager.HttpConnection
	MongoDBConnection mongodbmanager.MongoDBConfiguration
}

type Registrable interface {
	GetHttpRouterHandlers() []httpmanager.HttpRouterHandler
	GetEmptyData() interface {}

	GetId() int
	SetId( int )
}

//The manager object
type Manager struct {

	nextId int
	registered map[int]Registrable

	configuration ManagerContext
	httpManager httpmanager.HttpManager
	mongoDBManager mongodbmanager.MongoDBManager
}

//Register business objects
func (m *Manager) Construct( configuration ManagerContext ) {
	m.registered = make(map[int]Registrable)
	m.configuration = configuration
	//Register with the http manager so it listens to the correct endpoints
	m.httpManager.Construct( m.Execute )
	m.mongoDBManager.Construct( configuration.MongoDBConnection )

}

//Register business objects
func (m *Manager) Register( registrable Registrable ) {
	m.nextId++
	registrable.SetId( m.nextId )
	m.registered[m.nextId] = registrable

	//Register with the http manager so it listens to the correct endpoints
	m.httpManager.Register( registrable )
}

//Handler the start up of the manager
func (m *Manager) Start( ) {
	//start up server execution will wait here until the server is shutdown
	m.httpManager.Start( m.configuration.HttpConnection );
}

//This is the main call back method form all http requests
func (m *Manager) Execute( context httpmanager.HttpContext ) {
	fmt.Println("Executing", context)
	theData := m.registered[context.ProcessorId].GetEmptyData()

	//Handle payload if it exists
	payload, err := ioutil.ReadAll(context.Request.Body)

	fmt.Println( payload, ": ", len(payload), ": ", err )

	//There is no body if body is nil or EOF err is returned
	if payload != nil && len(payload) > 0 && err == nil {
		fmt.Println("Has Body")
		defer context.Request.Body.Close() //make sure we clean up the steam


		//
		err = json.Unmarshal(payload, &theData)
		if err != nil {
			http.Error(context.Writer, err.Error(), 500)
			fmt.Println(err)
			return
		}
	}


	for _, method := range context.RouteHandler.EndPointMethods {
		if method.HttpMethod == context.Request.Method {
			result := m.mongoDBManager.Execute( method.Callback, theData )
			json.NewEncoder(context.Writer).Encode(result) //stream the encoded data on the writer

			break; //not point in continuing to loop
		}
	}
}


