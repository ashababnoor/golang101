package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Capture connection properties
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "172.17.0.2:3306",
		DBName: "recordings",
	}
	// get a dabase handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	// db, err := sql.Open("mysql", "root:anypass@tcp(172.17.0.2:3306)/recordings")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")
}
