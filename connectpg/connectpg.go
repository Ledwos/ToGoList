package connectpg

// connection to database and query fxns for routes

import (
	// "fmt"
	"net/http"
	"database/sql"
	"log"
	"os"

	// routes "github.com/Ledwos/ToGoList/routing"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func Dbconnect() {
	// load env data
	loadEnv := godotenv.Load()
	if loadEnv != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to db
	db, err := sql.Open("postgres", os.Getenv("DB_STRING"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// call ping
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	
}

var loadEnv = godotenv.Load()
var db, err = sql.Open("postgres", os.Getenv("DB_STRING"))

func AddTask(c *gin.Context) {

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
}

func GetAll(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM usr")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			// panic(err)
			c.JSON(http.StatusOK, gin.H {
				"Error": "no rows returned :(",
			})
		}
		c.JSON(http.StatusOK, gin.H {
			"Success": "Got Data",
			"ID": id,
			"Name": name,
			"Age": age,
		})
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
//  ^ inefficient call
