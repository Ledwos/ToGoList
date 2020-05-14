package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	// when ready to be deployed, set to ReleaseMode
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//serve static files (html / css / js)
	router.Use(static.Serve("/", static.LocalFile("./", true)))

	// test route
	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "You woke me up %s", name)
	})

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
	router.Run(":"+os.Getenv("PORT"))
}
