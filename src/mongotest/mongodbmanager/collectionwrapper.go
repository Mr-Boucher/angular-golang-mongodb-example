package mongodbmanager

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"bytes"
)

type CollectionWrapper interface {
	Find(query map[string]string) *mgo.Query
	Insert(docs ...interface{}) error
	Remove(selector interface{}) error
	Update(selector interface{}, update interface{}) error
	GetQueryString() string
}

//
type collectionObj struct {
	collection *mgo.Collection
	queryString string
}

//Constructor
func NewCollectionWrapper( collection *mgo.Collection ) CollectionWrapper {
	return &collectionObj{collection:collection}
}

//getters
func (c *collectionObj) GetQueryString() string{
	return c.queryString
}


//
func createKeyValuePairs(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}

// With returns a copy of c that uses session s.
//func (c *Collection) With(s *mgo.Session) *mgo.Collection {
//	result := c.collection.With( s )
//	return result
//}
//
//func (c *Collection) EnsureIndexKey(key ...string) error {
//	result := c.collection.EnsureIndex(mgo.Index{Key: key})
//	return result
//}
//
//func (c *Collection) EnsureIndex(index mgo.Index) error {
//	result := c.collection
//	return result
//}

//func (c *Collection) DropIndex(key ...string) error {
//	result := c.collection
//	return result
//}
//
//func (c *Collection) DropIndexName(name string) error {
//	result := c.collection
//	return result
//}
//
//func (c *Collection) Indexes() (indexes []mgo.Index, err error) {
//	result := c.collection
//	return result
//}

func (c *collectionObj) Find(query map[string]string) *mgo.Query {
	result := c.collection.Find( bson.M{} )
	queryData := make(map[string]string)
	queryData["Query Type"] = "Find"
	for key, value := range query {
		queryData[key] = value
	}

	c.queryString = createKeyValuePairs( queryData )
	return result
}

//func (c *Collection) Repair() *mgo.Iter {
//	result := c.collection
//	return result
//}
//
//func (c *Collection) FindId(id interface{}) *mgo.Query {
//	result := c.collection
//	return result
//}
//
//func (c *Collection) Pipe(pipeline interface{}) *mgo.Pipe {
//	result := c.collection
//	return result
//}
//
//func (c *Collection) NewIter(session *mgo.Session, firstBatch []bson.Raw, cursorId int64, err error) *mgo.Iter {
//	result := c.collection
//	return result
//}

func (c *collectionObj) Insert(docs ...interface{}) error {
 	fmt.Println( "CollectionWrapper::Overriden Insert" )
	result := c.collection.Insert( docs[0] ) //take the 0th element for no obvious reason
	return result
}

func (c *collectionObj) Update(selector interface{}, update interface{}) error {
	result := c.collection.Update( selector, update )
	return result
}
//
//func (c *Collection) UpdateId(id interface{}, update interface{}) error {
//	result := c.collection.UpdateId( id, update )
//	return result
//}
//
//func (c *Collection) UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
//	result, err := c.collection.UpdateAll( selector, update )
//	return result, err
//}
//
//func (c *Collection) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
//	result, err	 := c.collection.Upsert( selector, update )
//	return result, err
//}
//
//func (c *Collection) UpsertId(id interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
//	result, err := c.collection.UpsertId( id, update )
//	return result, err
//}
//
func (c *collectionObj) Remove(selector interface{}) error {
	result := c.collection.Remove(selector)
	return result
}
//
//func (c *Collection) RemoveId(id interface{}) error {
//	result := c.collection.RemoveId( id )
//	return result
//}

//func (c *Collection) RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error) {
//	result := c.collection
//	return result
//}

//func (c *Collection) DropCollection() error {
//	result := c.collection
//	return result
//}
//
//func (c *Collection) Create(info *mgo.CollectionInfo) error {
//	result := c.collection.Create( info )
//	return result
//}

//func (c *Collection) Count() (n int, err error) {
//	result, err := c.collection.Count()
//	return result, err
//}

//func (c *Collection) writeOp(op interface{}, ordered bool) (lerr *mgo.LastError, err error) {
//	result := c.collection
//	return result
//}
//
//func (c *Collection) writeOpQuery(socket *mongoSocket, safeOp *mgo.QueryOp, op interface{}, ordered bool) (lerr *mgo.LastError, err error) {
//	result := c.collection
//	return result
//}
//
//func (c *Collection) writeOpCommand(socket *mongoSocket, safeOp *mgo.QueryOp, op interface{}, ordered, bypassValidation bool) (lerr *mgo.LastError, err error) {
//	result := c.collection
//	return result
//}



