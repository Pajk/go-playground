package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pajk/go-playground/simple-api/config"
	"github.com/pajk/go-playground/simple-api/routes"
	"github.com/pajk/go-playground/simple-api/utils"
)

func main() {
	utils.Init(config.App.MongoURL, config.App.MongoDB)

	r := gin.Default()

	r.GET("/say/:something", routes.Say)
	r.GET("/", routes.Hello)
	r.POST("/data", routes.Data)
	r.POST("/data-db", routes.PostDataDB)
	r.GET("/goroutines", routes.Goroutines)

	fmt.Printf("Server starting on :%s\n", config.App.Port)

	r.Run(":" + config.App.Port)
}
