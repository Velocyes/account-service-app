package config

import (
	"fmt"
	"log"
	"os"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

var DB *sql.DB

func ConnectToDB() {
	user := os.Getenv("user")
	pass := os.Getenv("pass")
	host := os.Getenv("host")
	port := os.Getenv("port")
	dbname := os.Getenv("dbname")
	
	var connectionString = user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbname

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("error open connection", err.Error())
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("error connect to db", errPing.Error())
	} else {
		fmt.Println("berhasil")
	}
	DB = db
}
