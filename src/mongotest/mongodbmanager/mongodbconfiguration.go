package mongodbmanager

//
type MongoDBConfiguration struct {
	DatabaseName   string
	CollectionName string
	Cluster        []string
	UserDatabase   string
	Username       string
	Password       string
	ConnectionTimeOut int
}