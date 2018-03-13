package httpmanager

//
type HttpMethodFunction interface {
	GetHttpMethod() string
	GetCallback() func( context interface{}, arguments interface{} ) interface{}
}

//
type httpMethodFunctionObject struct {
	httpMethod string
	callback func( context interface{}, arguments interface{} ) interface{}
}

//
func NewHttpMethodFunction( httpMethod string, callback func( context interface{}, arguments interface{} ) interface{} ) HttpMethodFunction {
	return &httpMethodFunctionObject{httpMethod, callback}
}

//
func (h *httpMethodFunctionObject) GetHttpMethod() string {
	return h.httpMethod
}

//
func (h *httpMethodFunctionObject) GetCallback() (func( context interface{}, arguments interface{} ) interface{}) {
	return h.callback
}