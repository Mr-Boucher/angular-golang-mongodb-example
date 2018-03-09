package manager

import (
	"net/http"
	"../httpmanager"
	"fmt"
	"../mongodbmanager"
)

type ManagerContext struct {
	httpConnection httpmanager.HttpConnection
	mongoDBConnection mongodbmanager.MongoDBConfiguration
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
	m.httpManager.Initialize( )
}

//Register business objects
func (m *Manager) Register( registrable Registrable ) {
	//Register with the http manager so it listens to the correct endpoints
	m.httpManager.Register( registrable )
}

//Handler the start up of the manager
func (m *Manager) Start( context ManagerContext ) {
	//start up server execution will wait here until the server is shutdown
	m.httpManager.Start( context.httpConnection, m.Execute );
}

//This is the main call back method form all http requests
func (m *Manager) Execute( writer http.ResponseWriter, request *http.Request ) {
	fmt.Println( "Executing" )
}


