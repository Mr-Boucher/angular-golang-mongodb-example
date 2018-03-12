package applicationmanager

import (
	"../httpmanager"
	"../mongodbmanager"
)

//
type ApplicationConfiguration struct {
	httpConnection httpmanager.HttpConnection
	mongoDBConfiguration mongodbmanager.MongoDBConfiguration
}

//Factory Method
func NewApplicationConfiguration( httpConnection httpmanager.HttpConnection, mongoDBDefault mongodbmanager.MongoDBConfiguration ) *ApplicationConfiguration {
	return &ApplicationConfiguration{httpConnection, mongoDBDefault}
}
