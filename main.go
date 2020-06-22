package main

import (
	// "os"

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
	// router.Use(static.Serve("/", static.LocalFile("./tglfront/build", true)))
	//cors - to tie it all together
	router.Use(cors.Default())

	//call route handler
	routes.Routes(router)

	// start / run server on given port
	// router.Run(":"+os.Getenv("PORT"))
	// CHANGE router.Run AND UNCOMMENT OS IMPORT WHEN DEPLOYING! - DEVELOPMENT ONLY
	router.Run(":8080")
}