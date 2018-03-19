package mongodbmanager

import (
	"gopkg.in/mgo.v2"
	"fmt"
)


type ContextHolder interface {
	GetMongoDBContext() MongoDBContext
}

//
type MongoDBContext interface {
	GetCollection() *Collection
	GetSession() *mgo.Session
	GetConfiguration() *MongoDBConfiguration

	SetCollection( *Collection )
	SetSession( *mgo.Session )
	SetConfiguration( *MongoDBConfiguration )
}

//
type MongoDBContextObject struct {
	configuration *MongoDBConfiguration
	session       *mgo.Session
	collection    *Collection
}

func (co *MongoDBContextObject ) GetCollection() *Collection {
	return co.collection
}

func (co *MongoDBContextObject ) SetCollection( collection *Collection ) {
	co.collection = collection
}

func (co *MongoDBContextObject ) GetSession() *mgo.Session {
	return co.session
}

func (co *MongoDBContextObject ) SetSession( session *mgo.Session ) {
	co.session = session
}

func (co *MongoDBContextObject ) GetConfiguration() *MongoDBConfiguration {
	return co.configuration
}

func (co *MongoDBContextObject ) SetConfiguration(configuration *MongoDBConfiguration ) {
	fmt.Println( "MongoDBContext::SetConfiguration to", configuration )
	co.configuration = configuration
}

