package data

import (
	"database/sql"
	"fmt"
)

//Db - our database
var Db *sql.DB

//Init - Initilizate
func Init() {
	var err error
	Db, err = sql.Open("sqlite3", "forum")
	if err != nil && Db == Db {
		fmt.Println(err)
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	return
}

//CloseDB - to close db
func CloseDB() {
	Db.Close()
	return
}
