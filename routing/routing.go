package routing

import (
	"net/http"
	// "database/sql"

	dbcon "github.com/Ledwos/ToGoList/connectpg"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)


func Routes(router *gin.Engine) {

	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "You woke me up %s", name)
	})

	// route group for pg calls 
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {
				"message": "pong",
			})
		})
	}

	//add api routes below
	api.POST("/add/:name/:age", dbcon.AddTask)
	// api.GET("/get/:user", dbcon.GetUser)
	api.GET("/better", dbcon.GetBetter)
	api.POST("/body", dbcon.Readbody)
	api.POST("/newacc", dbcon.Newacc)
	api.GET("/login", dbcon.Loguserin)

}
