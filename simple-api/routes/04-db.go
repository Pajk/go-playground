package routes

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/pajk/go-playground/simple-api/utils"
)

// TequestData Incomming request
type RequestData struct {
	Title   string `json:"title" binding:"required"`
	BrandID int64  `json:"brand_id" binding:"required"`
}

// Listing data from DB
type Listing struct {
	ID      string `json:"id" bson:"_id"`
	Title   string `json:"title"`
	BrandID int64  `json:"brand_id"`
}

// PostDataDB handles requests to /data-db route
func PostDataDB(c *gin.Context) {
	var req RequestData

	if c.BindJSON(&req) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data",
		})
		return
	}

	db := utils.Collection("listings")

	err := db.Insert(req)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to store the data",
		})
		return
	}

	fromDB := Listing{}

	err = db.Find(bson.M{"title": req.Title}).One(&fromDB)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to fetch from database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": fromDB,
	})
}
