package connectpg

// connection to database and query fxns for routes

import (
	// "fmt"
	// "encoding/json"
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

// struct version
func GetBetter(c *gin.Context) {
	// query string
	sqlStat := `SELECT * FROM usr`

	// query db
	rows, err := db.Query(sqlStat)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	type rec struct {
		Id int			
		Name string		
		Age int			
	}

	result := []rec{}

	// build response and return JSON
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"Error": "no rows returned :(",
			})
		}
		row := rec{id, name, age}
		result = append(result, row)
	}
	c.JSON(http.StatusOK, result)

}
