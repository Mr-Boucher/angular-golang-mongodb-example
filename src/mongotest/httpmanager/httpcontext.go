package httpmanager

import (
	"net/http"
)
//
type HttpContext struct {
	ProcessorId int
	Writer http.ResponseWriter
	Request *http.Request
	RouteHandler HttpRouterHandler
	Params map[string]string
}

//
type HttpConnection struct {
	Port int
}