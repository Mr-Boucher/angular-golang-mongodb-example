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
	httpManager httpmanager.HttpManager
}

//Register business objects
func (m *Manager) Initialize() {
	//Register with the http manager so it listens to the correct endpoints
	m.httpManager.Initialize( m.Execute)
}

//Register business objects
func (m *Manager) Register( registrable Registrable ) {
	//Register with the http manager so it listens to the correct endpoints
	m.httpManager.Register( registrable )
}

//Handler the start up of the manager
func (m *Manager) Start( context ManagerContext ) {
	//start up server execution will wait here until the server is shutdown
	m.httpManager.Start( context.HttpConnection );
}

//This is the main call back method form all http requests
func (m *Manager) Execute( context httpmanager.HttpContext ) {
	fmt.Println("Executing", context)
	for _, method := range context.RouteHandler.EndPointMethods {
		if method.HttpMethod == context.Request.Method {
			result := method.Callback( )
			json.NewEncoder(context.Writer).Encode(result) //stream the encoded data on the writer
		}
	}
}


