package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB = nil

func ConnectToDB() {
	user := os.Getenv("user")
	pass := os.Getenv("pass")
	host := os.Getenv("host")
	port := os.Getenv("port")
	dbname := os.Getenv("dbname")

	connectionString := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, errOpen := sql.Open("mysql", connectionString)
	if errOpen != nil {
		log.Fatal("Error open connection", errOpen.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	errPing := db.Ping()
	if errPing != nil {
		DB = nil
		log.Fatal("Error when connecting to db", errPing.Error())
	} else {
		DB = db
	}
}
