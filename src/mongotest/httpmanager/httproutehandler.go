package httpmanager

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

//
type HttpRouterHandler interface {
	GetURL() string
	GetEndPointMethods() []HttpMethodFunction
	ServeHTTP(http.ResponseWriter, *http.Request)
}

//Created by
type httpRouterHandlerObject struct {
	processorId     int
	url             string
	endPointMethods []HttpMethodFunction
	processor       Processor
}

//
func NewHttpRouteHandler(ProcessorId int, URL string, EndPointMethods []HttpMethodFunction) HttpRouterHandler {
	h := httpRouterHandlerObject{ProcessorId, URL, EndPointMethods, nil}
	return &h
}

//
func (h *httpRouterHandlerObject) GetURL() string {
	return h.url
}

//
func (h *httpRouterHandlerObject) GetEndPointMethods() []HttpMethodFunction {
	return h.endPointMethods
}

//
func (handler *httpRouterHandlerObject) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var err error

	fmt.Println("HttpRouterHandler:ServeHTTP Method", request.Method)
	fmt.Println("HttpRouterHandler:ServeHTTP Path", request.URL.Path)
	fmt.Println("HttpRouterHandler:ServeHTTP Params", mux.Vars(request))

	request.ParseForm() //make sure the values are filled in otherwise request.Form and request.PostForm are empty
	for key, value := range request.Form {
		fmt.Println("HttpRouterHandler:ServeHTTP Form ", key, "=", value )
	}
	for key, value := range request.PostForm {
		fmt.Println("HttpRouterHandler:ServeHTTP PostForm ", key, "=", value )
	}

	//Create the execution context
	context := HttpContext{handler.processorId, writer, request, handler, mux.Vars(request), request.Form }
	handler.setHeaders(context)

	//Call the action function requests
	if context.Request.Method != "OPTIONS" {
		err = handler.processor.Execute(context)
		if err != nil {
			fmt.Println("HttpRouterHandler::The Error:", err)
			http.Error(writer, err.Error(), 500)
		}
	}

	fmt.Println("HttpRouterHandler:Finished", request.Method)
}

//Set headers to tell the client what is supported for this REST API
func (h *httpRouterHandlerObject) setHeaders(context HttpContext) {
	var methodsStrings = "OPTIONS, HEAD, "
	for _, endpoint := range context.RouteHandler.GetEndPointMethods() {
		methodsStrings += endpoint.GetHttpMethod() + ","
	}

	fmt.Println("MethodsStrings: ", methodsStrings)
	context.Writer.Header().Add("Access-Control-Allow-Origin", "*")                                            //Allow access from anywhere
	context.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Request-Origin") //Allows setting of the Content-Type by the client
	context.Writer.Header().Add("Access-Control-Allow-Methods", methodsStrings)       //REST API supports GET, POST, PUT, DELETE
	context.Writer.Header().Add("Accept", "application/json")                                                  //Only json is accepted
	context.Writer.Header().Add("Content-Type", "application/json")
}