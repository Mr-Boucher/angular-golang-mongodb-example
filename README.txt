# Initial Setup
* install git and SourceTree
* Install Go https://golang.org/dl/
* https://nodejs.org/en/
* Open a terminal and run "npm install -g @angular/cli"

# IntelliJ setup
* install Go plugin 
* set GOROOT to your go install directory
* Get all go dependencies. Open a termainal and run "go get .." in the folder that contains your Go files.

You can now run the project either in a docker container or locally. 

# Use the following to build as a docker image (working dir is the project root)
docker build -t golang-test src/mongotest

# Use the following to run as a docker container
docker run -p 8000:8000 -h golang-server golang-test

# To run outside docker container
Go to main.go file and run the "main" function. In a terminal go to client/src and run "npm install" and "ng serve"
