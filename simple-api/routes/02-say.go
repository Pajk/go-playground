package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Say handles requests to /say route
func Say(c *gin.Context) {
	something := c.Param("something")

	c.String(http.StatusOK, "Hello from Go! %s", something)
}
