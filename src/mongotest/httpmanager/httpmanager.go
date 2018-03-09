package httpmanager

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
	"reflect"
	"encoding/json"
	"../configuration"
	"strings"
)

//The manager
type HttpManager struct {
	router mux.Router
  supported []NeedsHTTPSupport
  routingMap map[string]HttpRouterHandler
}


//
type HttpContext struct {
	writer http.ResponseWriter
	request *http.Request
}

//
type HttpHandler interface {
	Add(handler HttpRouterHandler)
}

//
type NeedsHTTPSupport interface {
	InitializeHTTPSupport()
	GetURL() string
	GetHTTPMethods() []string
}


//
type HttpConnection struct {
	port int
}

//
type HttpMethodFunction struct {
	httpMethod string
	function   *func()
}

//
type HttpRouterHandler struct {
	URL             string
	EndPointMethods []HttpMethodFunction
	ActionFunction  *func( processcontext.ExecutionContext ) interface{}
}

//
type Registrable interface {
	GetHttpRouterHandlers() []HttpRouterHandler
}

//
func (m *HttpManager) Initialize() {
	//
	m.router = mux.NewRouter(); //create the underlying http router
	m.supported = make([]NeedsHTTPSupport, 0); //create the empty default list of supported
	m.routingMap = make(map[string]HttpRouterHandler); //create the empty default list of router handlers
}

//
func (m *HttpManager) Register(registrable Registrable) {

	//Loop through all the routes
	for _, handler := range registrable.GetHttpRouterHandlers() {
		if handler.URL == nil {
			panic( "Route is nil" )
		}

		//make sure it is not already registered
		if m.routingMap[handler.URL] != nil {
			panic( "Route " + handler.URL + " is already defined" )
		}

		//Cache the handler for use in the execute when the request comes in
		m.routingMap[handler.URL] = handler;

		//Create endpoint methods for all handler endpoints
		for _, endPoint := range handler.EndPointMethods {
			m.router.HandleFunc(handler.URL, httpExecute).Methods( endPoint.function )
		}
	}
}

//
func (m *HttpManager) Start(httpConnection HttpConnection, requestCallback func( http.ResponseWriter, *http.Request ) ) {

	//initial all http supported object
	for _, supportedItem := range m.supported {

		//Create list of HTTPMethods for this endpoint
		var temp []HttpMethodFunction
		for _, httpMethod := range supportedItem.GetHTTPMethods() {
			append( temp, HttpMethodFunction{httpMethod, requestCallback} )
		}
		Add(HttpRouterHandler{supportedItem.GetURL(), temp })
		supportedItem.InitializeHTTPSupport()
	}

	//Setup route for incoming data requests

	//router.HandleFunc("/data", options).Methods("OPTIONS")             //Setup data the REST API and call options
	//router.HandleFunc("/data", getData).Methods("GET")                   //Setup data as the REST API and call GetData for get requests
	//router.HandleFunc("/data", createData).Methods("POST")               //Setup data the REST API and call CreateData for delete requests

	//router.HandleFunc("/data/{id:[a-z0-9]+}", options).Methods("OPTIONS") //Setup data the REST API and options
	//router.HandleFunc("/data/{id:[a-z0-9]+}", updateData).Methods("PUT")    //Setup data the REST API and call UpdateData for delete requests
	//router.HandleFunc("/data/{id:[a-z0-9]+}", deleteData).Methods("DELETE") //Setup data the REST API and call DeleteData for delete requests

	//Start listening for requests - thread waits forever at this port
	log.Fatal(http.ListenAndServe(":" + httpConnection.port, router))
}

//Set headers to tell the client what is supported for this REST API
func setHeaders( context processcontext.ExecutionContext ) {
	var methodsString = "OPTIONS, HEAD"
	methodsString += strings.Join( context.HttpRouterHandler.EndPointMethods, ", ")

	context.HttpContext.writer.Header().Add("Access-Control-Allow-Origin", "*")                                            //Allow access from anywhere
	context.HttpContext.writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Request-Origin") //Allows setting of the Content-Type by the client
	context.HttpContext.writer.Header().Add("Access-Control-Allow-Methods", methodsString)       //REST API supports GET, POST, PUT, DELETE
	context.HttpContext.writer.Header().Add("Accept", "application/json")                                                  //Only json is accepted
	context.HttpContext.writer.Header().Add("Content-Type", "application/json")
}

//This is the main algorithm for processing requests
func httpExecute( writer http.ResponseWriter, request *http.Request ) {

	//Scrub the url if needed
	fmt.Println( "Recieved Route URL" + request.URL.Path )
	url := request.URL.Path

	//Get HttpRouterHandler
	routeHandler := routingMap[url]
	if routeHandler == nil { panic( "No route handler for URL " + url ) }

	//Create the execution context
	context := processcontext.ExecutionContext{configuration.MongoDB, HttpContext{writer, request}, routeHandler}

	//Call the action function requests
	setHeaders( context )
	if request.Method != "OPTIONS" {
		fmt.Println("Executing", reflect.TypeOf(routeHandler.ActionFunction))
		result := routeHandler.ActionFunction(context)
		json.NewEncoder(writer).Encode(result) //stream the encoded data on the writer
	}
}