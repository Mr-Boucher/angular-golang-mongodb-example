package main

import (
	"fmt"
	"strings"
	"os"
	"./httphandler"
	"./dataeditor"
)

//Kick it all off
func main() {
	//Log start up arguments
	fmt.Println(strings.Join(os.Args, " "))

	//add http handler objects
	httphandler.AddNeedsHTTPSupport( dataeditor.DataEditor{}() )

	//start up server execution will wait here until the server is shutdown
  httphandler.Start( httphandler.HttpConnection{8000} );
}






