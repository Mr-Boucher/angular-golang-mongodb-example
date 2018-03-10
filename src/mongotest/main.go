package main

import (
	"fmt"
	"strings"
	"os"
	"./dataeditor"
	"./applicationmanager"
	"./httpmanager"
	"./mongodbmanager"
)

//Should be a constant but can't because of language restriction that const can't have arrays
var mongoDBDefault = mongodbmanager.MongoDBConfiguration{"dev", "test", []string{"cluster0-shard-00-00-iaz9w.mongodb.net:27017"}, "admin", "dev", "dev"}
var httpConnectionDefault = httpmanager.HttpConnection{8000}

//Kick it all off
func main() {
	//Log start up arguments
	fmt.Println(strings.Join(os.Args, " "))

	//Create the manger that handles everything
	manager := applicationmanager.Manager{}

	//create the context
	configuration := applicationmanager.ManagerContext{ httpConnectionDefault, mongoDBDefault }

	//
	manager.Initialize( configuration )

	//Register business logic processors with the manager
	manager.Register( &dataeditor.DataEditor{} )
	//manager.Register( configuration.MongoDBConfiguration{}() )

	//Kick everything off
	manager.Start( )
}






