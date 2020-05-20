package routing

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// _ "github.com/lib/pq"
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

}



//add api routes below
// api.POST("/add/:name/:age", func(c *gin.Context) {
// 	name := c.Param("name")
// 	age := c.Param("age")
// 	sqlInsert := `
// 		INSERT INTO usr (name, age)
// 		VALUES ($1, $2)
// 		RETURNING name`
// 		id := ""
// 		err = db.QueryRow(sqlInsert, name, age).Scan(&id)
// 		if err != nil {
// 			panic(err)
// 		}
// 	c.JSON(http.StatusOK, gin.H {
// 			"Success": "new user added!",
// 	})
// })

// api.GET("/all", func(c *gin.Context) {
// 	rows, err := db.Query("SELECT * FROM usr")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var id int
// 		var name string
// 		var age int
// 		err = rows.Scan(&id, &name, &age)
// 		if err != nil {
// 			// panic(err)
// 			c.JSON(http.StatusOK, gin.H {
// 				"Error": "no rows returned :(",
// 			})
// 		}
// 		c.JSON(http.StatusOK, gin.H {
// 			"Success": "Got Data",
// 			"ID": id,
// 			"Name": name,
// 			"Age": age,
// 		})
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		panic(err)
// 	}
// })
// ^ inefficient call, store locally and then return JSON




