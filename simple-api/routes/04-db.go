package routes

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/pajk/go-playground/simple-api/utils"
)

// Listing data from DB
type Listing struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Title   string        `json:"title" bson:"title" binding:"required"`
	BrandID int64         `json:"brand_id" bson:"brand_id" binding:"required"`
}

// GetListing returns a listing from the database
func GetListing(c *gin.Context) {
	id := c.Param("id")
	fmt.Printf("%s", id)
	collection := utils.Collection("listings")
	var result Listing

	err := collection.FindId(bson.ObjectIdHex(id)).One(&result)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetListings returns all listings from the database
func GetListings(c *gin.Context) {
	collection := utils.Collection("listings")
	var result []Listing

	err := collection.Find(nil).All(&result)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

// CreateListing perists a new listing in the database
func CreateListing(c *gin.Context) {
	var listing Listing

	if c.BindJSON(&listing) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data",
		})
		return
	}

	db := utils.Collection("listings")

	listing.ID = bson.NewObjectId()

	err := db.Insert(listing)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to store the data",
		})
		return
	}

	c.JSON(http.StatusOK, listing)
}
