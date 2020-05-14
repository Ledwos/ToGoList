package main

import (
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	//serve static files (html / css / js)
	router.Use(static.Serve("/", static.LocalFile("./", true)))

	// route group for API 
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {
				"message": "pong",
			})
		})
	}

	//add api routes below


	// start / run server on given port
	router.Run()
}