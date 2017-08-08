package utils

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session
var database *mgo.Database

// Init connects to the database
func Init(url string, db string) {
	fmt.Println(url, db)

	var err error
	session, err = mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	database = session.DB(db)
}

// Close the database connection
func Close() {
	session.Close()
	session = nil
	database = nil
}

// Collection return db collection
func Collection(collection string) *mgo.Collection {
	return database.C(collection)
}
