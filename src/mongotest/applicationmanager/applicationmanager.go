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
type ApplicationManager interface {
	Register( registrable Registrable )
	Start( )
}

//The manager object
type applicationManagerObject struct {
	nextId int
	registered map[int]Registrable
	configuration *ApplicationConfiguration
	httpManager httpmanager.HttpManager
	mongoDBManager *mongodbmanager.MongoDBManager
}

//Register business objects
func NewApplicationManager( configuration ApplicationConfiguration ) ApplicationManager {
	fmt.Println( "ApplicationManager::ApplicationConfiguration", configuration  )
	m := applicationManagerObject{}
	m.registered = make(map[int]Registrable)
	m.configuration = &configuration
	//Register with the http manager so it listens to the correct endpoints
	m.httpManager = httpmanager.NewHttpManager( &m )
	m.mongoDBManager = mongodbmanager.NewMongoDBManager( configuration.mongoDBConfiguration )
	fmt.Println( "ApplicationManager::NewApplicationManager", m  )
	return &m
}

//Register business objects
func (m *applicationManagerObject) Register( registrable Registrable ) {
	m.nextId++
	registrable.SetId( m.nextId )
	m.registered[m.nextId] = registrable

	//Register with the http manager so it listens to the correct endpoints
	m.httpManager.Register( registrable )
	fmt.Println( "ApplicationManager::Register", registrable  )
}

//Handler the start up of the manager
func (m *applicationManagerObject) Start( ) {
	//start up server execution will wait here until the server is shutdown
	fmt.Println( "ApplicationManager::Start"  )
	m.httpManager.Start( m.configuration.httpConnection );
}

//This is the main call back method form all http requests
func (m *applicationManagerObject) Execute( httpcontext httpmanager.HttpContext ) {
	fmt.Println("ApplicationManager::Execute httpContext", httpcontext)

	//
	fmt.Println("ApplicationManager::Execute ApplicationConfiguration:", m.configuration)

	//Create the application context needed by all the managers
	context := ApplicationContext{}
	context.configuration = m.configuration
	context.httpContext = &httpcontext
	context.mongoDBContext = &mongodbmanager.MongoDBContextObject{}
	context.parameters = httpcontext.Params

	fmt.Println("ApplicationManager::Execute GetConfiguration:", context.mongoDBContext.GetConfiguration())
	fmt.Println("ApplicationManager::Execute SetConfiguration:", context.configuration.mongoDBConfiguration)
	context.mongoDBContext.SetConfiguration( &context.configuration.mongoDBConfiguration )
	fmt.Println("ApplicationManager::Execute GetConfiguration:", context.mongoDBContext.GetConfiguration())
	fmt.Println("ApplicationManager::Execute MongoContext:", context.GetMongoDBContext() )
	fmt.Println("ApplicationManager::Execute MongoConfiguration:", context.mongoDBContext.GetConfiguration() )

	//use a factory method to get the processors specific data object
	fmt.Println("ApplicationManager::Execute Process", m.registered[context.httpContext.ProcessorId])

	//Handle payload if it exists
	payload, err := ioutil.ReadAll(context.httpContext.Request.Body)
	fmt.Println( payload, ": ", len(payload), ": ", err )

	//There is no body if body is nil or EOF err is returned
	var theData interface{}
	if payload != nil && len(payload) > 0 && err == nil {
		fmt.Println("ApplicationManager::Execute Has Body:", payload )
		defer context.httpContext.Request.Body.Close() //make sure we clean up the steam

		//
		theData, err = m.registered[context.httpContext.ProcessorId].Unmarshal( payload )
		fmt.Println("ApplicationManager::Execute TheData:", theData )

		if err != nil {
			http.Error(context.httpContext.Writer, err.Error(), 500)
			fmt.Println(err)
			return
		}
	}

	//find the correct end point
	//TODO make this a map instead of list
	fmt.Println("ApplicationManager::Execute Finding router:", context )
	for _, method := range context.httpContext.RouteHandler.GetEndPointMethods() {
		if method.GetHttpMethod() == context.httpContext.Request.Method {
			fmt.Println("ApplicationManager::Execute Context:", context )
			fmt.Println("ApplicationManager::Execute MongoContext:", context.GetMongoDBContext() )
			//Make sure the Mongo context is connected to the mongo configuration
			m.mongoDBManager.InitContext( &context )
			defer m.mongoDBManager.CleanupContext( &context ) //make sure the context for mongo is cleaned up

			//call data store
			result := m.mongoDBManager.Execute( &context, method.GetCallback(), theData )

			if result != nil {
				byteData, err := json.Marshal(result)
				if err != nil {
					panic(err)
				}

				context.httpContext.Writer.Write(byteData)
			}

			break; //no point in continuing to loop it found the method
		}
	}
}


