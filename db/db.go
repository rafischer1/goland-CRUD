package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//db commands
	db, err := sql.Open("mysql",
		"golangTest")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Ping!")
	}
	defer db.Close()
}
