package httpmanager

//
type Registrable interface {
	GetHttpRouterHandlers() []HttpRouterHandler
}

//
type registered interface {
	GetURL() string
	GetHTTPMethods() []string
}
