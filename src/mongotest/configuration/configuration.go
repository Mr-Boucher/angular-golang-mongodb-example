package configuration

type MongoDBConfiguration struct {
	DatabaseName   string
	CollectionName string
	Cluster        []string
	UserDatabase   string
	Username       string
	Password       string
}

//Should be a constant but can't because of language restriction that const can't have arrays
var MongoDB = MongoDBConfiguration{"dev", "test", []string{"cluster0-shard-00-00-iaz9w.mongodb.net:27017"}, "admin", "dev", "dev"}



