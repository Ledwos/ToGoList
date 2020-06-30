package connectpg

// connection to database and query fxns for routes

import (
	"net/http"
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
	// Dev only
	"log"
	"github.com/joho/godotenv"

)

func Dbconnect() {
	// Dev only
	loadEnv := godotenv.Load()
	if loadEnv != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to db
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
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

// Dev only
var loadEnv = godotenv.Load()

var db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))

// create new user
func Newacc(c *gin.Context) {
	// form data structure
	type Accform struct {
		Name string `json:"name" binding:"required"`
		Pass string `json:"pass" binding:"required"`
		Email string `json:"email" binding:"required"`
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
		Email string `json:"email" binding:"required"`
		Pass string `json:"pass" binding:"required"`
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
		panic(err)
		// c.JSON(http.StatusBadRequest, gin.H {
		// 	"error": err,
		// })
	} else {
		c.JSON(http.StatusOK, gin.H {
			"username": rtn_name,
			"userid": rtn_id,
		})
	}
}

// // add new task for user
func NewTask(c *gin.Context) {
	// task form
	type taskForm struct {
		Userid   string	`json:"userid"   binding:"required"`
		Taskname string `json:"taskname" binding:"required"`
		Taskdesc string `json:"taskdesc" binding:"required"`
		Taskdate string `json:"taskdate" binding:"required"`
		Tasktime string `json:"tasktime" binding:"required"`
	}
	var json taskForm
	c.Bind(&json)
	// get form data
	taskname := json.Taskname
	taskdesc := json.Taskdesc
	taskdate := json.Taskdate
	tasktime := json.Tasktime
	userid := json.Userid

	//queries for all cases

	// name only
	if (taskdesc == "none" && taskdate == "none" && tasktime == "none") {
		sqlTask := `
		INSERT INTO t_table (t_name)
		VALUES ($1)
		RETURNING t_id
		`
		t_id := ""
		err := db.QueryRow(sqlTask, taskname).Scan(&t_id)
		if err != nil {
			panic(err)
		}
		ut_res := newTask2(t_id, userid)
		c.JSON(http.StatusOK, gin.H {
			"Success": "task added!",
			"task id": ut_res,
		})
		return
	// no desc / date
	} else if (taskdesc == "none" && taskdate == "none") {
		sqlTask := `
		INSERT INTO t_table (t_name, t_time)
		VALUES ($1, $2)
		RETURNING t_id
		`
		t_id := ""
		err := db.QueryRow(sqlTask, taskname, tasktime).Scan(&t_id)
		if err != nil {
			panic(err)
		}
		ut_res := newTask2(t_id, userid)
		c.JSON(http.StatusOK, gin.H {
			"Success": "task added!",
			"task id": ut_res,
		})
		return
	// no desc / time
	} else if (taskdesc == "none" && tasktime == "none") {
		sqlTask := `
		INSERT INTO t_table (t_name, t_date)
		VALUES ($1, $2)
		RETURNING t_id
		`
		t_id := ""
		err := db.QueryRow(sqlTask, taskname, taskdate).Scan(&t_id)
		if err != nil {
			panic(err)
		}
		ut_res := newTask2(t_id, userid)
		c.JSON(http.StatusOK, gin.H {
			"Success": "task added!",
			"task id": ut_res,
		})
		return
	// no date / time
	} else if (taskdate == "none" && tasktime == "none") {
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
		return

	// no desc
	} else if (taskdesc == "none") {
		sqlTask := `
		INSERT INTO t_table (t_name, t_date, t_time)
		VALUES ($1, $2, $3)
		RETURNING t_id
		`
		t_id := ""
		err := db.QueryRow(sqlTask, taskname, taskdate, tasktime).Scan(&t_id)
		if err != nil {
			panic(err)
		}
		ut_res := newTask2(t_id, userid)
		c.JSON(http.StatusOK, gin.H {
			"Success": "task added!",
			"task id": ut_res,
		})
		return

	// no date
	} else if (taskdate == "none") {
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
		return
	// no time
	} else if (tasktime == "none") {
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
		return
	// all fields completed
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
		return
	}
}



// ALTERNATE ADD A USER 
// add new task for user
// func NewTask(c *gin.Context) {
// 	// task json data
// 	type taskForm struct {
// 		Userid   string	`json:"userid" binding:"required"`
// 		Taskname string `json:"taskname" binding:"required"`
// 		Taskdesc string `json:"taskdesc" binding:"required"`
// 		Taskdate string `json:"taskdate" binding:"required"`
// 		Tasktime string `json:"tasktime" binding:"required"`
// 	}
// 	var json taskForm
// 	c.Bind(&json)
// 	// get form data
// 	userid := json.Userid
// 	taskname := json.Taskname
// 	taskdesc := json.Taskdesc
// 	taskdate := json.Taskdate
// 	tasktime := json.Tasktime

// 	type res struct {
// 		Uid   string
// 		Tname string		
// 		Tdesc string		
// 		Tdate string
// 		Ttime string
// 	}

// 	result := []res{}

// 	empty := " "
// 	if taskdesc == " " {
// 		empty = "none"
// 	} else {
// 		empty = taskdesc
// 	}

// 	row := res{userid, taskname, empty, taskdate, tasktime}
// 	result = append(result, row)

// 		// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//         // c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
//         // c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
//         // c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
// 	c.JSON(http.StatusOK, result)
// 	return
// }


// sql query to update bridging table UNCOMMENT
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
		SELECT t_id, t_name, t_desc, t_date, t_time
		FROM ut_table ut
		INNER JOIN t_table t ON t.t_id = ut.utt_id
		WHERE t_comp = false AND ut.utu_id = $1
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
		// Tcomp bool
	}

	result := []res{}

	// build response and return JSON
	for rows.Next() {
		var tid   int
		var tname string		
		var tdesc string		
		var tdate string
		var ttime string
		// var tcomp bool
		err = rows.Scan(&tid, &tname, &tdesc, &tdate, &ttime)
		row := res{tid, tname, tdesc, tdate, ttime}
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
		Tid 	int		`json:"taskid" binding:"required"`
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
		Taskid int `json:"taskid" binding:"required"`
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
		panic(err)
	} else {
		// c.Header("Access-Control-Allow-Origin", "*")
		// c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        // c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        // c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        // c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.JSON(http.StatusOK, gin.H {
			"response": 1,
		})
		return
	}
}

