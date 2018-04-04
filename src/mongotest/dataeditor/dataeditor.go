package dataeditor

import (
	"fmt"
	"reflect"
	"regexp"

	"gopkg.in/mgo.v2/bson"
	"github.com/rs/xid"

	"encoding/json"

	"../httpmanager"
	"../mongodbmanager"
	"../utils"
	"strconv"
)

const (
	baseUrl = "/data"
)

//
type Context interface {
	GetParameters() map[string]string
	GetSliceParameters() map[string][]string
	GetCollection() mongodbmanager.CollectionWrapper
}

//
type DataEditor interface {
	GetId() int
	SetId(int)
	GetHttpRouterHandlers() []httpmanager.HttpRouterHandler
	//Marshal( theData interface {} ) ([]byte, error)
	Unmarshal([]byte) (interface{}, error)
}

//
type dataEditorObject struct {
	id    int
	regex *regexp.Regexp
}

//
func NewEditor() DataEditor {
	regex, err := regexp.Compile("[^a-zA-Z\\d\\s]")

	fmt.Println("Error creating search validation regex", err)

	config := dataEditorObject{regex:regex}
	return &config
}

//GetHttpRouterHandlers() meets the interface
func (d *dataEditorObject) GetHttpRouterHandlers() []httpmanager.HttpRouterHandler {

	routers := []httpmanager.HttpRouterHandler{}

	//Create the call back methods for /Data
	dataMethodFunctions := []httpmanager.HttpMethodFunction{}
	dataMethodFunctions = append(dataMethodFunctions, httpmanager.NewHttpMethodFunction("GET", d.search))
	dataMethodFunctions = append(dataMethodFunctions, httpmanager.NewHttpMethodFunction("POST", d.create))
	dataMethodFunctions = append(dataMethodFunctions, httpmanager.NewHttpMethodFunction("DELETE", d.deleteAll))
	dataMethodHandler := httpmanager.NewHttpRouteHandler(d.id, baseUrl, dataMethodFunctions)
	routers = append(routers, dataMethodHandler)

	//Search
	dataSearchMethodFunctions := []httpmanager.HttpMethodFunction{}
	dataSearchMethodFunctions = append(dataSearchMethodFunctions, httpmanager.NewHttpMethodFunction("GET", d.search))
	dataSearchMethodHandler := httpmanager.NewHttpRouteHandler(d.id, baseUrl + "?pageNumber=[0-9]+&search=[a-zA-Z0-9]+", dataSearchMethodFunctions)
	routers = append(routers, dataSearchMethodHandler)

	//add the backs method for /Data/id
	dataIdMethodFunctions := []httpmanager.HttpMethodFunction{}
	dataIdMethodFunctions = append(dataIdMethodFunctions, httpmanager.NewHttpMethodFunction("PUT", d.update))
	dataIdMethodFunctions = append(dataIdMethodFunctions, httpmanager.NewHttpMethodFunction("DELETE", d.deleteById))
	dataIdMethodHandler := httpmanager.NewHttpRouteHandler(d.id, baseUrl + "/{id:[a-z0-9]+}", dataIdMethodFunctions)
	routers = append(routers, dataIdMethodHandler)

	return routers
}

//
func (d *dataEditorObject) Unmarshal(payload []byte) (interface{}, error) {

	theData := TestData{}
	err := json.Unmarshal(payload, &theData)

	fmt.Println("Data Editor:", theData)

	return theData, err
}

func (d *dataEditorObject) GetId() int {
	return d.id
}

func (d *dataEditorObject) SetId(id int) {
	d.id = id
}

//Load data from mongo returned as a []TestData
func (d *dataEditorObject) search(appcontext interface{}, arguments interface{}) (interface{}, error) {

	fmt.Println("DataEditor::Search", arguments)
	var results TestDataContainer
	var err error //default err value of nil
	var queryStr bson.M = bson.M{} //default criteria is to return everything
	var criteria searchCriteria

	context := appcontext.(Context)
	collection := context.GetCollection()

	//Get the search criteria
	criteria, err = d.parseCriteria( context )

	//
	if err == nil && len(criteria.filter) > 0 {
		//validate search criteria does not contain
		fmt.Println("DataEditor::Search searchCriteria", criteria.filter)
		hasInvalidSearchCriteria := d.regex.MatchString(criteria.filter)
		if hasInvalidSearchCriteria {
			err = utils.NewError(criteria.filter + " contains invalid chars")
		} else {
			fmt.Println("DataEditor::Search searchCriteria", criteria.filter)

			//Create bson regex expression from the search criteria
			regex := bson.RegEx{}
			searchType := "" //starts with
			regex.Pattern = searchType + criteria.filter //
			regex.Options = "i" //make search case-insensitive
			queryStr = bson.M{"value": regex }
		}
	}

	//Do the actual search in MongoDB if no errors
	if err == nil {
		query := collection.Find(queryStr) //
		query = query.Sort("value") //sort the data by its value
		query = query.Skip(criteria.pageSize * (criteria.pageNumber-1) ) //todo this is a known performance issue
		query = query.Limit(criteria.pageSize) //Limit the size after the sort
		query.All(&results.TestData) //execute the query

		//
		results.TotalCount,err = collection.Find(queryStr).Count()
		fmt.Println("results.TotalCount", results.TotalCount)

		//Display the data returned for debugging
		fmt.Println("***********************Start of results***********************")
		for index, result := range results.TestData {
			fmt.Println(index, "id:", result.Id, "value:", result.Value)
		}
		fmt.Println("***********************End of results***********************")
		fmt.Println("Finished Search data")
	}

	return results, err
}

type searchCriteria struct {
	filter string
	pageNumber int
	pageSize int
}

//Parse the criteria
func (d *dataEditorObject) parseCriteria(context Context) (searchCriteria, error) {
	var err error

	//parse and validate search criteria
	filter := context.GetSliceParameters()["search"][0]
	fmt.Println("DataEditor::Search searchCriteria", filter)

	//Empty value is fine, if not empty then validate
	if err == nil && len(filter) > 0 {
		//validate search criteria does not contain
		fmt.Println("DataEditor::Search searchCriteria", filter)
		hasInvalidSearchCriteria := d.regex.MatchString(filter)
		if hasInvalidSearchCriteria {
			err = utils.NewError(filter + " contains invalid chars")
		}
	}

	//parse and validate pageNumber
	var pageNumber int
	pageNumberStr := context.GetSliceParameters()["pageNumber"]
	fmt.Println("DataEditor::Search pageNumber", pageNumberStr)
	if err == nil && len(pageNumberStr) <= 0 {
		err = utils.NewError("Missing pageNumber")
	} else {
		pageNumber, err = strconv.Atoi(pageNumberStr[0])
		if pageNumber <= 0 {
			err = utils.NewError("Invalid pageNumber " + pageNumberStr[0])
		}
	}

	//parse and validate pageSize
	var pageSize int
	pageSizeStr := context.GetSliceParameters()["pageSize"]
	fmt.Println("DataEditor::Search pageSize", pageSizeStr, len(pageSizeStr))
	if err == nil && len(pageSizeStr) <= 0 {
		err = utils.NewError("Missing pageSize")
	} else {
		pageSize, err = strconv.Atoi(pageSizeStr[0])
		if pageSize <= 0 {
			err = utils.NewError("Invalid pageSize " + pageSizeStr[0])
		}
	}

	//
	if err != nil {
		fmt.Println("DataEditor::Search pageSize", err)
	}

	return searchCriteria{filter, pageNumber, pageSize}, err
}

////remove data from db base
func (d *dataEditorObject) create(appcontext interface{}, arguments interface{}) (interface{}, error) {
	fmt.Println("DataEditor::create arguments", arguments)
	var err error
	context := appcontext.(Context)

	//Validation handling
	if arguments == nil {
		panic("Missing argument of type TestData")
	}

	fmt.Println("arguments:", arguments)

	//Convert the empty interface to a string that contains the id
	newData, ok := arguments.(TestData) //same as casting in java
	if !ok {
		errorMessage := fmt.Sprint("Argument should be of type TestData. It was ", reflect.TypeOf(arguments))
		panic(errorMessage)
	}

	//make sure we create the id before storing it
	newData.Id = xid.New().String()

	//Insert the TestData
	fmt.Println("create:", newData)
	collection := context.GetCollection()
	err = collection.Insert(newData)

	//creating new one the first element is the only one that needs a new id
	fmt.Println("create:", "finished", newData)

	return newData, err
}

//remove data from db base
func (d *dataEditorObject) update(context interface{}, arguments interface{}) (interface{}, error) {
	fmt.Println("DataEditor::update arguments", arguments)
	var err error

	//id, ok := arguments.(string)
	//if !ok {
	//	panic( "Argument should be of type string" )
	//}
	//
	//collection := context.(*mgo.Collection)
	//collection.Remove( bson.M{"id": id} )
	fmt.Println("update:", "finished")
	return nil, err
}

//remove data from db base
func (d *dataEditorObject) deleteById(appcontext interface{}, arguments interface{}) (interface{}, error) {
	context := appcontext.(Context)
	var err error
	id := context.GetParameters()["id"] //get the id of the object to delete
	fmt.Println("DataEditor::deleteById", context.GetParameters(), "started")

	collection := context.GetCollection()
	fmt.Println("DataEditor::deleteById collection", collection)
	err = collection.Remove(bson.M{"id": id})

	fmt.Println("DataEditor::deleteById", id, "finished")
	return nil, err
}

//remove data from db base
func (d *dataEditorObject) deleteAll(appcontext interface{}, arguments interface{}) (interface{}, error) {

	fmt.Println("DataEditor::deleteAll arguments", arguments)
	var err error
	//context := appcontext.(Context)
	//id := context.GetParameters()["id"] //get the id of the object to delete
	//
	//fmt.Println( "arguments:", arguments )
	//
	////Validation handling
	//if arguments == nil {
	//	panic( "Missing argument of type string" )
	//}
	//
	////Convert the empty interface to a string that contains the id
	//id, ok := arguments.(string)
	//if !ok {
	//	panic( "Argument should be of type string" )
	//}
	//
	//fmt.Println("deleteById:", id, "started")
	//collection := context.GetCollection()
	//err := collection.Remove( bson.M{"id": id} )
	//if err != nil {
	//	panic( err )
	//}
	//fmt.Println("deleteById:", id, "finished")
	return nil, err
}
