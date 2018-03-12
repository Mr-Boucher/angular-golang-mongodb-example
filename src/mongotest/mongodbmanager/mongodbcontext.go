package mongodbmanager

import (
	"gopkg.in/mgo.v2"
)


type ContextHolder interface {
	GetMongoDBContext() MongoDBContext
}

//
type MongoDBContext interface {
	GetCollection() *mgo.Collection
	GetSession() *mgo.Session
	GetConfiguration() *MongoDBConfiguration

	SetCollection( *mgo.Collection )
	SetSession( *mgo.Session )
	SetConfiguration( *MongoDBConfiguration )
}

//
type MongoDBContextObject struct {
	configuration *MongoDBConfiguration
	session       *mgo.Session
	collection    *mgo.Collection
}

func (co MongoDBContextObject ) GetCollection() *mgo.Collection {
	return co.collection
}

func (co MongoDBContextObject ) GetSession() *mgo.Session {
	return co.session
}

func (co MongoDBContextObject ) GetConfiguration() *MongoDBConfiguration {
	return co.configuration
}

func (co MongoDBContextObject ) SetCollection( collection *mgo.Collection ) {
	co.collection = collection
}

func (co MongoDBContextObject ) SetSession( session *mgo.Session ) {
	co.session = session
}

func (co MongoDBContextObject ) SetConfiguration(configuration *MongoDBConfiguration ) {
	co.configuration = configuration
}

