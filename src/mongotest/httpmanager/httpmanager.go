package httpmanager

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
	"reflect"
	"strconv"
)

//
type HttpManager interface {
	Register(registrable Registrable)
	Start( httpConnection HttpConnection )
}

//The manager
type httpManagerObject struct {
	routingMap []HttpRouterHandler
	router *mux.Router
  registered []registered
	processor Processor
}

//
func NewHttpManager( processor Processor ) HttpManager {
	//
	m := httpManagerObject{}
	m.router = mux.NewRouter(); //create the underlying http router
	m.registered = make([]registered, 0); //create the empty default list of supported
	m.routingMap = []HttpRouterHandler{}; //create the empty default list of router handlers
	m.processor = processor;
	fmt.Println( "HttpManager::NewHttpManager", m )
	return &m
}

//
func (m *httpManagerObject) Register(registrable Registrable) {

	fmt.Println( "HttpManager::Register: " + reflect.TypeOf(registrable).String())

	//Loop through all the routes
	for _, handler := range registrable.GetHttpRouterHandlers() {

		handler.(*httpRouterHandlerObject).processor = m.processor //set the application main algorithm
		fmt.Println( "HttpManager::Registering handler: " + reflect.TypeOf(handler).String())

		//Cache the handler for use in the execute when the request comes in
		m.routingMap = append( m.routingMap, handler )

		//Create endpoint methods for all handler endpoints
		url := handler.GetURL()
		m.router.Handle(url, handler).Methods( "OPTIONS" ) //options must always be set for CORS to work
		for _, endPoint := range handler.GetEndPointMethods() {
			fmt.Println( "HttpManager::Registering endpoint:", url, endPoint )
			m.router.Handle(url, handler).Methods( endPoint.GetHttpMethod() )
		}
	}
}

//
func (m *httpManagerObject) Start( httpConnection HttpConnection )  {

	//Start listening for requests - thread waits forever at this port
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(httpConnection.Port), m.router))
}


