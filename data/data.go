package data

import (
	"database/sql"
	"log"
)

//Db - our database
var Db *sql.DB

//Init - Initilizate
func Init() {
	var err error
	Db, err = sql.Open("sqlite3", "forum")
	defer Db.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = Db.Ping()
	if err != nil {
		log.Fatal("Init: failed to connect database")
	}

}
