package applicationmanager

import (
	"../httpmanager"
	"../mongodbmanager"
)

//
type ApplicationConfiguration interface {
	GetHttpConnection() httpmanager.HttpConnection
	GetMongoDBConfiguration() mongodbmanager.MongoDBConfiguration
}

type applicationConfiguration struct {
	httpConnection httpmanager.HttpConnection
	mongoDBConfiguration mongodbmanager.MongoDBConfiguration
}

//Factory Method
func NewApplicationConfiguration( httpConnection httpmanager.HttpConnection, mongoDBDefault mongodbmanager.MongoDBConfiguration ) ApplicationConfiguration {
	return &applicationConfiguration{httpConnection:httpConnection, mongoDBConfiguration:mongoDBDefault}
}

func (a *applicationConfiguration) GetHttpConnection() httpmanager.HttpConnection {
	return a.httpConnection
}

func (a *applicationConfiguration) SetHttpConnection( connection httpmanager.HttpConnection ) {
	a.httpConnection = connection
}


func (a *applicationConfiguration) GetMongoDBConfiguration() mongodbmanager.MongoDBConfiguration {
	return a.mongoDBConfiguration
}

func (a *applicationConfiguration) SetDBConfiguration( mongoDBConfiguration mongodbmanager.MongoDBConfiguration ) {
	a.mongoDBConfiguration = mongoDBConfiguration
}
