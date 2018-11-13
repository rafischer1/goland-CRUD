package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("listening on :8080")
	fmt.Fprintf(w, "Golang CRUD APP, %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	//db commands
	db, err := sql.Open("sql",
		"golangTest")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
