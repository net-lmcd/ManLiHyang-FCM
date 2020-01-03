package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const localConfig = "root:tkdals12@tcp(localhost:3306)/ManLiHyang"

var db *DB

func initDatabase() {
	db, err := sql.Open("mysql", localConfig)
	if err != nil {
		log.Fatal("[MYSQL DATABASE CONNECTION EXCEPTION : ", err.Error())
		panic(err.Error())
	}

	defer db.Close()

	// Open doesn't open a connection , Validate DSN data
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	log.Print("[DATABASE CONNECTION ESTABLISHED]")
}

func findTokenByUsn(usn string) []string {
	stmtOut, err := db.Prepare("SELECT token FROM tokens WHERE usn = ?")
	if err != nil {
		log.Print("[DATABASE PREPARE FAILED] : ", err.Error())
		return nil
	}

	var userUsn string
	err = stmtOut.QueryRow(usn).Scan(&userUsn)
	if err != nil {
		log.Print("[DATABASE OPERATION FAILED] : ", err.Error())
		return nil
	}
	return []string{userUsn}
}

func saveToken(usn string, token string) bool {
	stmtIns, err := db.Prepare("INSERT INTO tokens VALUES( ?, ? )") // ? = placeholder
	if err != nil {
		log.Print("[DATABASE PREPARE FAILED] : ", err.Error())
		return false
	}

	_, err = stmtIns.Exec(token, usn)
	if err != nil {
		log.Print("[DATABASE OPERATION FAILED] : ", err.Error())
		return false
	}
	return true
}
