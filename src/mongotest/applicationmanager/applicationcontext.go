package applicationmanager

import(
	"../httpmanager"
	"../mongodbmanager"
)

//
type ContextHolder interface {
	GetMongoDBContext() mongodbmanager.MongoDBContext
	GetHttpContext() *httpmanager.HttpContext
	GetParameters() map[string]string
	GetSliceParameters() map[string][]string
	GetCollection() mongodbmanager.CollectionWrapper
}

//
type applicationContext struct {
	configuration ApplicationConfiguration
	parameters map[string]string
	sliceParameters map[string][]string
	httpContext *httpmanager.HttpContext
	mongoDBContext mongodbmanager.MongoDBContext
}

//
func (ac *applicationContext) GetMongoDBContext() mongodbmanager.MongoDBContext {
	return ac.mongoDBContext
}

//
func (ac *applicationContext) GetHttpContext() *httpmanager.HttpContext {
	return ac.httpContext
}

//
func (ac *applicationContext) GetParameters() map[string]string {
	return ac.parameters
}

//
func (ac *applicationContext) GetSliceParameters() map[string][]string {
	return ac.sliceParameters
}

//
func (ac *applicationContext) GetConfiguration() ApplicationConfiguration {
	return ac.configuration
}

//
func (ac *applicationContext) GetCollection() mongodbmanager.CollectionWrapper {
	return ac.mongoDBContext.GetCollection()
}