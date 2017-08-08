package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hello handles requests to / route
func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "From Go!",
	})
}
