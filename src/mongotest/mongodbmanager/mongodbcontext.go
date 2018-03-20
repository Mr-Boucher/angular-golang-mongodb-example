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
	GetCollection() CollectionWrapper
	GetSession() *mgo.Session
	GetConfiguration() *MongoDBConfiguration

	SetCollection( CollectionWrapper )
	SetSession( *mgo.Session )
	SetConfiguration( *MongoDBConfiguration )
}

//
type mongoDBContextObject struct {
	configuration *MongoDBConfiguration
	session       *mgo.Session
	collection    CollectionWrapper
}

//Constructor
func NewMongoDBContext() MongoDBContext {
	return &mongoDBContextObject{}
}

func (co *mongoDBContextObject ) GetCollection() CollectionWrapper {
	return co.collection
}

func (co *mongoDBContextObject ) SetCollection( collection CollectionWrapper ) {
	co.collection = collection
}

func (co *mongoDBContextObject ) GetSession() *mgo.Session {
	return co.session
}

func (co *mongoDBContextObject ) SetSession( session *mgo.Session ) {
	co.session = session
}

func (co *mongoDBContextObject ) GetConfiguration() *MongoDBConfiguration {
	return co.configuration
}

func (co *mongoDBContextObject ) SetConfiguration(configuration *MongoDBConfiguration ) {
	fmt.Println( "MongoDBContext::SetConfiguration to", configuration )
	co.configuration = configuration
}

