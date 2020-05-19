package main

import (
	"net/http"
	// "os"
	"fmt"
	"database/sql"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	dbname = "test_db"
)

// Why are these not read by pq?
// var user = os.Getenv("PG_USER")
// var password = os.Getenv("PG_PASS")

func main() {

	// when ready to be deployed, set to ReleaseMode
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//serve static files (html / css / js)
	router.Use(static.Serve("/", static.LocalFile("./", true)))

	// db connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	
	// connect to db
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// call ping
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//test Insert SQL statement
	// sqlInsert := `
	// INSERT INTO usr (name, age)
	// VALUES ($1, $2)
	// RETURNING name`
	// name := ""
	// err = db.QueryRow(sqlInsert, "Gopher", "11").Scan(&name)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(name, " has been added to the database!")

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
	api.POST("/add/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		sqlInsert := `
			INSERT INTO usr (name, age)
			VALUES ($1, $2)
			RETURNING name`
			id := ""
			err = db.QueryRow(sqlInsert, name, age).Scan(&id)
			if err != nil {
				panic(err)
			}
		c.JSON(http.StatusOK, gin.H {
				"Success": "new user added!",
		})
	})

	api.GET("/all", func(c *gin.Context) {
		type Allusr struct {
			ID		int
			Name 	string
			Age 	int
		}
		sqlAll := `SELECT * FROM usr;`
		var user Allusr
		row := db.QueryRow(sqlAll)
		err := row.Scan(&user.ID, &user.Name, &user.Age)
		switch err {
		case sql.ErrNoRows:
			c.JSON(http.StatusOK, gin.H {
				"Error": "no rows returned :(",
			})
			return
		case nil:
			c.JSON(http.StatusOK, gin.H {
				"Success": user,
			})
		default:
			panic(err)
		}
	})

	// start / run server on given port
	// router.Run(":"+os.Getenv("PORT"))
	// CHANGE router.Run AND UNCOMMENT OS IMPORT WHEN DEPLOYING!
	router.Run()
}