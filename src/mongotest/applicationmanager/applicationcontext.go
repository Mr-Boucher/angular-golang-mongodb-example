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
	GetCollection() *mongodbmanager.Collection
}

//
type ApplicationContext struct {
	configuration *ApplicationConfiguration
	parameters map[string]string
	httpContext *httpmanager.HttpContext
	mongoDBContext mongodbmanager.MongoDBContext
}

//
func (ac *ApplicationContext) GetMongoDBContext() mongodbmanager.MongoDBContext {
	return ac.mongoDBContext
}

//
func (ac *ApplicationContext) GetHttpContext() *httpmanager.HttpContext {
	return ac.httpContext
}

//
func (ac *ApplicationContext) GetParameters() map[string]string {
	return ac.parameters
}

//
func (ac *ApplicationContext) GetConfiguration() *ApplicationConfiguration {
	return ac.configuration
}

//
func (ac *ApplicationContext) GetCollection() *mongodbmanager.Collection {
	return ac.mongoDBContext.GetCollection()
}