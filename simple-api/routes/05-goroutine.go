package routes

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/pajk/go-playground/simple-api/utils"
)

func fetchDB(done chan string) {
	var fromDB Listing
	db := utils.Collection("listings")
	err := db.Find(bson.M{"title": "Yeezy"}).One(&fromDB)
	if err != nil {
		fmt.Println("Fetch error", err.Error())
	}
	fmt.Println(fromDB)
	done <- fromDB.Title
}

func fetchURL(url string, done chan string) {
	done <- url
}

// Goroutines handler
func Goroutines(c *gin.Context) {

	messages := make(chan string, 2)

	go fetchDB(messages)
	go fetchURL("http://www.google.com", messages)

	msg1 := <-messages
	msg2 := <-messages

	c.String(http.StatusOK, "DONE msg1: %s, msg2: %s", msg1, msg2)
}
