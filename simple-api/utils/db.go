package utils

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session
var database *mgo.Database

// InitDB connects to the database
func InitDB(url string, db string) {
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
	if database == nil {
		panic("Database connection not configured")
	}
	return database.C(collection)
}
