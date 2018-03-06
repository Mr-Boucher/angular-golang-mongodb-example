package httphandler

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//
type HttpHandler interface {
	Add(handler HttpRouterHandler)
}

//
type NeedsHTTPSupport interface {
	initializeHTTPSupport()
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
	url             string
	endPoints       []HttpMethodFunction
	actionFunction  *func(http.ResponseWriter, *http.Request)
	optionsFunction *func(http.ResponseWriter, *http.Request)
}

//
var router = mux.NewRouter(); //create the underlying http router
var routerHandlers []HttpRouterHandler = make([]HttpRouterHandler, 0); //create the empty default list of router handlers
var supported []NeedsHTTPSupport = make([]NeedsHTTPSupport, 0); //create the empty default list of supported

//
func AddSupported(handler HttpRouterHandler) {
	append(routerHandlers, handler) //record keep for the handlers
	router.HandleFunc("url", handler.optionsFunction).Methods("OPTIONS") //add security options

	//create endpoint methods for all handler endpoints
	for _, endPoint := range handler.endPoints {
		router.HandleFunc(handler.url, handler.actionFunction).Methods( endPoint.function ) //add security options
	}
}

//
func AddNeedsHTTPSupport(needsHTTPSupport NeedsHTTPSupport) {
	append( supported, needsHTTPSupport )
}

//
func Start(httpConnection HttpConnection) {

	//initial all http supported object
	for _, supportedItem := range supported {
		supportedItem.initializeHTTPSupport()
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
func setHeaders(writer http.ResponseWriter) {
	writer.Header().Add("Access-Control-Allow-Origin", "*")                                            //Allow access from anywhere
	writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Request-Origin") //Allows setting of the Content-Type by the client
	writer.Header().Add("Access-Control-Allow-Methods", "HEAD, GET, POST, PUT, DELETE, OPTIONS")       //REST API supports GET, POST, PUT, DELETE
	writer.Header().Add("Accept", "application/json")                                                  //Only json is accepted
	writer.Header().Add("Content-Type", "application/json")
}