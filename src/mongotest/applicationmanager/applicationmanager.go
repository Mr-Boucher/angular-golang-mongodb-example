package applicationmanager

import (
	"../httpmanager"
	"fmt"
	"../mongodbmanager"
	"encoding/json"
)

//
type ManagerContext struct {
	HttpConnection httpmanager.HttpConnection
	MongoDBConnection mongodbmanager.MongoDBConfiguration
}

type Registrable interface {
	GetHttpRouterHandlers() []httpmanager.HttpRouterHandler
}

//The manager object
type Manager struct {
	configuration ManagerContext
	httpManager httpmanager.HttpManager
	mongoDBManager mongodbmanager.MongoDBManager
}

//Register business objects
func (m *Manager) Initialize( configuration ManagerContext ) {
	m.configuration = configuration
	//Register with the http manager so it listens to the correct endpoints
	m.httpManager.Initialize( m.Execute )
	m.mongoDBManager.Initialize( configuration.MongoDBConnection )

}

//Register business objects
func (m *Manager) Register( registrable Registrable ) {
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
	for _, method := range context.RouteHandler.EndPointMethods {
		if method.HttpMethod == context.Request.Method {
			result := m.mongoDBManager.Execute( method.Callback, nil )
			json.NewEncoder(context.Writer).Encode(result) //stream the encoded data on the writer
		}
	}
}


