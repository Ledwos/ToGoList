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
		SELECT u_name, u_id FROM u_table
		WHERE u_email = $1 AND u_pass = $2
	`
	rtn_name := ""
	rtn_id := ""
	err := db.QueryRow(sqlLog, logEmail, logPass).Scan(&rtn_name, &rtn_id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H {
			"Error": "Make sure you type your email / password correctly",
			"Details": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H {
			"Welcome back": rtn_name,
			"User id": rtn_id,
		})
	}
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

func NewTask(c *gin.Context) {
	// task form
	type taskForm struct {
		Userid   string	`form:"userid" binding:"required"`
		Taskname string `form:"taskname" binding:"required"`
		Taskdesc string `form:"taskdesc" binding:"required"`
		Taskdate int 	`form:"taskdate" binding:"required"`
		Tasktime int 	`form:"tasktime" binding:"required"`
	}
	var json taskForm
	c.Bind(&json)
	// get form data
	taskname := json.Taskname
	taskdesc := json.Taskdesc
	taskdate := json.Taskdate
	tasktime := json.Tasktime
	userid := json.Userid

	// multiple queries in case time / date are not sent over
	if (taskdate == 0 && tasktime == 0) {
		sqlTask := `
		INSERT INTO t_table (t_name, t_desc)
		VALUES ($1, $2)
		RETURNING t_id
		`
		t_id := ""
		err := db.QueryRow(sqlTask, taskname, taskdesc).Scan(&t_id)
		if err != nil {
			panic(err)
		}
		ut_res := newTask2(t_id, userid)
		c.JSON(http.StatusOK, gin.H {
			"Success": "task added!",
			"task id": ut_res,
		})
	} else if (taskdate == 0) {
		sqlTask := `
		INSERT INTO t_table (t_name, t_desc, t_time)
		VALUES ($1, $2, $3)
		RETURNING t_id
		`
		t_id := ""
		err := db.QueryRow(sqlTask, taskname, taskdesc, tasktime).Scan(&t_id)
		if err != nil {
			panic(err)
		}
		ut_res := newTask2(t_id, userid)
		c.JSON(http.StatusOK, gin.H {
			"Success": "task added!",
			"task id": ut_res,
		})
	} else if (tasktime == 0) {
		sqlTask := `
		INSERT INTO t_table (t_name, t_desc, t_date)
		VALUES ($1, $2, $3)
		RETURNING t_id
		`
		t_id := ""
		err := db.QueryRow(sqlTask, taskname, taskdesc, taskdate).Scan(&t_id)
		if err != nil {
			panic(err)
		}
		ut_res := newTask2(t_id, userid)
		c.JSON(http.StatusOK, gin.H {
			"Success": "task added!",
			"task id": ut_res,
		})
	} else {
		sqlTask := `
		INSERT INTO t_table (t_name, t_desc, t_date, t_time)
		VALUES ($1, $2, $3, $4)
		RETURNING t_id
		`
		t_id := ""
		err := db.QueryRow(sqlTask, taskname, taskdesc, taskdate, tasktime).Scan(&t_id)
		if err != nil {
			panic(err)
		}
		ut_res := newTask2(t_id, userid)
		c.JSON(http.StatusOK, gin.H {
			"Success": "task added!",
			"task id": ut_res,
		})
	}

	// sql query to update bridging table
	// sqlMany := `
	// 	INSERT INTO ut_table (utu_id, utt_id)
	// 	VALUES ($1, $2)
	// 	RETURNING ut_id
	// `
	// ut_id := ""
	// err2 := db.QueryRow(sqlMany, userid, t_id).Scan(&ut_id)
	// if err2 != nil {
	// 	panic(err2)
	// } else {
	// 	c.JSON(http.StatusOK, gin.H {
	// 		"Success": "task added!",
	// 		"task id": ut_id,
	// })}

}

func newTask2(t_id, u_id string) string{
	// sql query to update bridging table
	sqlMany := `
		INSERT INTO ut_table (utu_id, utt_id)
		VALUES ($1, $2)
		RETURNING ut_id
	`
	ut_id := ""
	err2 := db.QueryRow(sqlMany, u_id, t_id).Scan(&ut_id)
	if err2 != nil {
		panic(err2)
	} else {
		return ut_id
	}
}







































// BELOW: for testing only

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

func GetTask(c *gin.Context) {
	// query string
	sqlStat := `SELECT * FROM t_table`

	// query db
	rows, err := db.Query(sqlStat)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	type rec struct {
		Id int			
		Name string		
		Desc string
		Created string
		Date string
		Time string
		Comp string
	}

	result := []rec{}

	// build response and return JSON
	for rows.Next() {
		var id int			
		var name string		
		var desc string
		var created string
		var date string
		var time string
		var comp string
		err = rows.Scan(&id, &name, &desc, &created, &date, &time, &comp)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"Error": "no rows returned :(",
			})
		}
		row := rec{id, name, desc, created, comp, date, time}
		result = append(result, row)
	}
	c.JSON(http.StatusOK, result)

}

func GetBridge(c *gin.Context) {
	// query string
	sqlStat := `SELECT * FROM t_table`

	// query db
	rows, err := db.Query(sqlStat)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	type rec struct {
		Uid int
		Tid int
		Utid int			
		
	}

	result := []rec{}

	// build response and return JSON
	for rows.Next() {
		var uid int
		var tid int
		var utid int
		err = rows.Scan(&uid, &tid, &utid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"Error": "no rows returned :(",
			})
		}
		row := rec{uid, tid, utid}
		result = append(result, row)
	}
	c.JSON(http.StatusOK, result)

}
