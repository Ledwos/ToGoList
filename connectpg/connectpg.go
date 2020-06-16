package connectpg

// connection to database and query fxns for routes

import (
	"net/http"
	"database/sql"
	"log"
	"os"

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

// create new user
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

// user authorisation
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

// add new task for user
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
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.JSON(http.StatusOK, gin.H {
			"Success": "task added!",
			"task id": ut_res,
		})
		return
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
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.JSON(http.StatusOK, gin.H {
			"Success": "task added!",
			"task id": ut_res,
		})
		return
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
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.JSON(http.StatusOK, gin.H {
			"Success": "task added!",
			"task id": ut_res,
		})
		return
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
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.JSON(http.StatusOK, gin.H {
			"Success": "task added!",
			"task id": ut_res,
		})
		return
	}
}

// sql query to update bridging table
func newTask2(t_id, u_id string) string{
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

// Change u_id back to body instead of url?
// Get tasks for given user
func GetTasks(c *gin.Context) {
	u_id := c.Param("id")
	// sql query to get all tasks for user u_id
	sqlTasks := `
		SELECT t_id, t_name, t_desc, t_date, t_time, t_comp
		FROM ut_table ut
		INNER JOIN t_table t ON t.t_id = ut.utt_id
		WHERE ut.utu_id = $1
	`
	// query db
	rows, err := db.Query(sqlTasks, u_id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	type res struct {
		Tid   int
		Tname string		
		Tdesc string		
		Tdate string
		Ttime string
		Tcomp bool
	}

	result := []res{}

	// build response and return JSON
	for rows.Next() {
		var tid   int
		var tname string		
		var tdesc string		
		var tdate string
		var ttime string
		var tcomp bool
		err = rows.Scan(&tid, &tname, &tdesc, &tdate, &ttime, &tcomp)
		row := res{tid, tname, tdesc, tdate, ttime, tcomp}
		result = append(result, row)
	}
	c.Header("Access-Control-Allow-Origin", "*")
    // c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, &result)

}

// change complete status of task (probably some on-click fxn)

func CompTask(c *gin.Context) {
	// task complete form
	type Comp struct {
		Tid 	int		`form:"taskid" binding:"required"`
	}
	var json Comp
	c.Bind(&json)
	// get form data
	taskid := json.Tid
	// sql query
	sqlComp := `
		UPDATE t_table
		SET t_comp = NOT t_comp
		WHERE t_id = $1 
	`
	_, err := db.Exec(sqlComp, taskid)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, gin.H {
			"Response": 1,
		})
	}
}

// update task name
func UpdateName(c *gin.Context) {
	// task form
	type taskForm struct {
		Taskid	 int	`form:"taskid" binding:"required"`
		Taskname string `form:"taskname" binding:"required"`
	}
	var json taskForm
	c.Bind(&json)
	// get form data
	taskid := json.Taskid
	taskname := json.Taskname

	sqlTask := `UPDATE t_table
		SET t_name = $2
		WHERE t_id = $1
	`
	_, err := db.Exec(sqlTask, taskid, taskname)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, gin.H {
			"Response": 1,
		})
	}
}

// update task description
func UpdateDesc(c *gin.Context) {
	// task form
	type taskForm struct {
		Taskid	 int	`form:"taskid" binding:"required"`
		Taskdesc string `form:"taskdesc" binding:"required"`
	}
	var json taskForm
	c.Bind(&json)
	// get form data
	taskid := json.Taskid
	taskdesc := json.Taskdesc
	
	sqlTask := `UPDATE t_table
		SET t_desc = $2
		WHERE t_id = $1
	`
	_, err := db.Exec(sqlTask, taskid, taskdesc)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, gin.H {
			"Response": 1,
		})
	}
}

// update task date
func UpdateDate(c *gin.Context) {
	// task form
	type taskForm struct {
		Taskid	 int	`form:"taskid" binding:"required"`
		Taskdate int 	`form:"taskdate" binding:"required"`
	}
	var json taskForm
	c.Bind(&json)
	// get form data
	taskid := json.Taskid
	taskdate := json.Taskdate
	
	sqlTask := `UPDATE t_table
		SET t_date = $2
		WHERE t_id = $1
	`
	_, err := db.Exec(sqlTask, taskid, taskdate)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, gin.H {
			"Response": 1,
		})
	}
}

// update task time
func UpdateTime(c *gin.Context) {
	// task form
	type taskForm struct {
		Taskid	 int	`form:"taskid" binding:"required"`
		Tasktime string `form:"tasktime" binding:"required"`
	}
	var json taskForm
	c.Bind(&json)
	// get form data
	taskid := json.Taskid
	tasktime := json.Tasktime
	
	sqlTask := `UPDATE t_table
		SET t_time = $2
		WHERE t_id = $1
	`
	_, err := db.Exec(sqlTask, taskid, tasktime)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, gin.H {
			"Response": 1,
		})
	}
}


// delete task
func DeleteTask(c *gin.Context) {
	// get task id
	type taskDel struct {
		Taskid int `form:"taskid" binding:"required"`
	}
	var json taskDel
	c.Bind(&json)
	taskid := json.Taskid
	// delete query
	sqlDel := `
		DELETE FROM t_table
		WHERE t_id = $1
	`
	_, err := db.Exec(sqlDel, taskid)
	if err != nil {
		// panic(err)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		c.JSON(http.StatusOK, gin.H {
			"response": "Something went wrong",
		})
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		c.JSON(http.StatusOK, gin.H {
			"response": 1,
		})
	}
}

