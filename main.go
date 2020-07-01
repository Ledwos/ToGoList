package main

import (
	// Deploy only
	"os"

	routes "github.com/Ledwos/ToGoList/routing"
	dbcon "github.com/Ledwos/ToGoList/connectpg"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	_ "github.com/lib/pq"
)

func main() {
	//connect to db
	dbcon.Dbconnect()

	// when ready to be deployed, set to ReleaseMode
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//serve static files (html / css / js)
	router.Use(static.Serve("/", static.LocalFile("./tglfront/build", true)))
	
	// cooperate with react router
	router.NoRoute(func(c *gin.Context) {
		c.File("./tglfront/build/index.html")
	})
	
	//cors - to tie it all together
	router.Use(cors.Default())

	//call route handler
	routes.Routes(router)

	// Deploy only
	// start / run server on given port 
	router.Run(":"+os.Getenv("PORT"))
	
	// Dev only
	// router.Run(":8080")
}