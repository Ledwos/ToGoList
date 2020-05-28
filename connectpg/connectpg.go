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
var db, err = sql.Open("postgres", os.Getenv("HDB_STRING"))

func Newacc(c *gin.Context) {
	// form data structure
	type Accform struct {
		Name string `form:"name" binding:"required"`
		Pass string `form:"pass" binding:"required"`
		Email string `form:"email" binding:"required"`
	}
	var json Accform
	c.Bind(&json)
	// get form data
	accName := json.Name
	accPass := json.Pass
	accEmail := json.Email
	// sql query
	sqlNewAcc := `
		INSERT INTO u_table (u_name, u_pass, u_email)
		VALUES ($1, $2, $3)
		RETURNING u_name
		`
	rtn_name := ""
	err := db.QueryRow(sqlNewAcc, accName, accPass, accEmail).Scan(&rtn_name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H {
			"Error": "email already used",
			"Details": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H {
			"New user!": rtn_name,
		})
	}
}

func Loguserin(c *gin.Context) {
	// form data structure
	type Logform struct {
		Email string `form:"email" binding:"required"`
		Pass string `form:"pass" binding:"required"`
	}
	var json Logform
	c.Bind(&json)
	// get form data
	logEmail := json.Email
	logPass := json.Pass
	// sql query
	// sqlLog := `
	// 	SELECT EXISTS(
	// 	SELECT 1 FROM u_table 
	// 	WHERE u_email = $1 AND u_pass = $2
	// 	)`
	sqlLog := `
		SELECT u_name, u_created_on FROM u_table
		WHERE u_email = $1 AND u_pass = $2
	`
	rtn_name := ""
	rtn_email := ""
	err := db.QueryRow(sqlLog, logEmail, logPass).Scan(&rtn_name, &rtn_email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H {
			"Error": "Make sure you type your email / password correctly",
			"Details": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H {
			"Welcome back": rtn_name,
			"your email": rtn_email,
		})
	}
}

func Readbody(c *gin.Context) {
	type Output struct {
		Something string `form:"something" binding:"required"`
		Name string `form:"name" binding:"required"`
		Three int `form:"three" binding:"required"`
		Newname int `form:"newname" binding:"required"`
	}

	var json Output
	c.Bind(&json)
	// jname := json.Name
	jsome := json.Something
	c.JSON(http.StatusOK, gin.H{
		"Got name": json,
		"Got thing": jsome,
	})
}

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
	sqlStat := `SELECT * FROM u_table`

	// query db
	rows, err := db.Query(sqlStat)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	type rec struct {
		Id int			
		Name string		
		Pass string
		Created string
		Email string
	}

	result := []rec{}

	// build response and return JSON
	for rows.Next() {
		var id int
		var name string
		var pass string
		var created string
		var email string
		err = rows.Scan(&id, &name, &pass, &created, &email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"Error": "no rows returned :(",
			})
		}
		row := rec{id, name, pass, created, email}
		result = append(result, row)
	}
	c.JSON(http.StatusOK, result)

}
