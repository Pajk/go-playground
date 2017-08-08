package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DataRequest describes data request structure
type DataRequest struct {
	Title   string `json:"title" binding:"required"`
	BrandID int64  `json:"brand_id" binding:"required"`
}

// Data handles requests to /data route
func Data(c *gin.Context) {
	var req DataRequest
	// var err error

	if c.BindJSON(&req) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": req,
	})
}
