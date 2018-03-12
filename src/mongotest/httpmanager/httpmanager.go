package httpmanager

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
	"reflect"
	"strconv"
	"regexp"
)

//
type Registrable interface {
	GetHttpRouterHandlers() []HttpRouterHandler
}

//
type RequestCallback interface {

}

//
type registered interface {
	GetURL() string
	GetHTTPMethods() []string
}

//
type HttpConnection struct {
	Port int
}

//Created by
type HttpRouterHandler struct {
	ProcessorId int
	URL             string
	EndPointMethods []HttpMethodFunction
}

//
type HttpMethodFunction struct {
	HttpMethod string
	Callback func( context interface{}, arguments interface{} ) interface{}
}

//The manager
type HttpManager struct {
	routingMap []HttpRouterHandler
	router *mux.Router
  registered []registered
	callback func( HttpContext )
}

//
type HttpContext struct {
	ProcessorId int
	Writer http.ResponseWriter
	Request *http.Request
	RouteHandler HttpRouterHandler
}

//This is the main algorithm for processing requests
func (m *HttpManager) httpExecute( writer http.ResponseWriter, request *http.Request ) {
	fmt.Println( "Executing:httpExecute" )
	//Scrub the url if needed
	fmt.Println( "Recieved Route URL" + request.URL.Path )
	url := request.URL.Path

	//TODO change to map of lists using base url as key
	var routeHandler *HttpRouterHandler
	for _, routerItem := range m.routingMap {
		match, _ := regexp.MatchString(routerItem.URL, url)
		if match {
			routeHandler = &routerItem
		}
	}

	//Get HttpRouterHandler
	if routeHandler == nil {
		panic( "No route handler for URL " + url )
		return
	}

	//Create the execution context
	context := HttpContext{routeHandler.ProcessorId, writer, request, *routeHandler}
	m.setHeaders( context )

	//Call the action function requests
	if context.Request.Method != "OPTIONS" {
		m.callback(context)
	}
}

//
func (m *HttpManager) Construct( requestCallback func( HttpContext ) ) {
	//
	m.router = mux.NewRouter(); //create the underlying http router
	m.registered = make([]registered, 0); //create the empty default list of supported
	m.routingMap = []HttpRouterHandler{}; //create the empty default list of router handlers
	m.callback = requestCallback;
}

//
func (m *HttpManager) Register(registrable Registrable) {

	fmt.Println( "Registering: " + reflect.TypeOf(registrable).String())

	//Loop through all the routes
	for _, handler := range registrable.GetHttpRouterHandlers() {

		fmt.Println( "Registering handler: " + reflect.TypeOf(handler).String())

		//make sure it is not already registered
		//_, ok := m.routingMap[handler.URL]
		//if ok {
		//	panic( "Route " + handler.URL + " is already defined" )
		//}

		//Cache the handler for use in the execute when the request comes in
		m.routingMap = append( m.routingMap, handler )

		//Create endpoint methods for all handler endpoints
		m.router.HandleFunc(handler.URL, m.httpExecute).Methods( "OPTIONS" ) //options must always be set for CORS to work
		for _, endPoint := range handler.EndPointMethods {
			fmt.Println( "Registering endpoint:", handler.URL, endPoint )
			m.router.HandleFunc(handler.URL, m.httpExecute).Methods( endPoint.HttpMethod )
		}
	}
}

//
func (m *HttpManager) Start( httpConnection HttpConnection  )  {

	//Start listening for requests - thread waits forever at this port
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(httpConnection.Port), m.router))
}

//Set headers to tell the client what is supported for this REST API
func (m *HttpManager) setHeaders( context HttpContext ) {
	var methodsStrings = "OPTIONS, HEAD, "
	for _, endpoint := range context.RouteHandler.EndPointMethods {
		methodsStrings += endpoint.HttpMethod + ","
	}

	fmt.Println( "MethodsStrings: ", methodsStrings)
	context.Writer.Header().Add("Access-Control-Allow-Origin", "*")                                            //Allow access from anywhere
	context.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Request-Origin") //Allows setting of the Content-Type by the client
	context.Writer.Header().Add("Access-Control-Allow-Methods", methodsStrings)       //REST API supports GET, POST, PUT, DELETE
	context.Writer.Header().Add("Accept", "application/json")                                                  //Only json is accepted
	context.Writer.Header().Add("Content-Type", "application/json")
}


