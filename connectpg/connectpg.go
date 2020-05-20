package connectpg

// connection to database and query fxns for routes

import (
	"fmt"
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	dbname = "test_db"
	user = ""
	password = ""
)

func Dbconnect() {

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
	fmt.Println("connect baby!")
}