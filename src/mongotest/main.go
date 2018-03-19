package main

import (
	"fmt"
	"strings"
	"os"
	"./applicationmanager"
	"./httpmanager"
	"./mongodbmanager"

	"./dataeditor"
	"./configuration"
)

//Should be a constant but can't because of language restriction that const can't have arrays
var mongoDBDefault = mongodbmanager.MongoDBConfiguration{"dev", "test", []string{"cluster0-shard-00-00-iaz9w.mongodb.net:27017"}, "admin", "dev", "dev"}
var httpConnectionDefault = httpmanager.HttpConnection{8000}

//Kick it all off
func main() {
	//Log start up arguments
	fmt.Println(strings.Join(os.Args, " "))

	//create the context
	appconfig := applicationmanager.NewApplicationConfiguration( httpConnectionDefault, mongoDBDefault )

	//
	manager := applicationmanager.NewApplicationManager( appconfig )

	//Register business logic processors with the manager
	manager.Register( dataeditor.NewEditor() )
	manager.Register( configuration.NewConfiguration() )

	//Kick everything off
	manager.Start( )
}






