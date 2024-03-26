package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(dbString string) (*sql.DB, error) {

	db, err := sql.Open("mysql", dbString)

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func InitDB(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully initialized the database")
}
