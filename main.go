package main

import (
	"database/sql"
	"fmt"
	mux "github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"

	"github.com/rafischer1/goCRUD/handlers"
)

var db *sql.DB

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func main() {
	fmt.Println("In main")
	initDb()
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/books", handlers.GetAll).Methods("GET")
	router.HandleFunc("/books", handlers.PostBook).Methods("POST")
	router.HandleFunc("/books/:id", handlers.EditBook).Methods("PUT")
	router.HandleFunc("/books/:id", handlers.DeleteBook).Methods("DELETE")

	// http.Handle("/", http.FileServer(http.Dir("./static")))
	// router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	router.Handle("/", http.FileServer(http.Dir("./static/")))
	router.Handle("static/", http.StripPrefix("static/", http.FileServer(http.Dir("static"))))

	log.Println("Listening...")
	http.ListenAndServe(":3000", router)
}

func initDb() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}
	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name
	return conf
}
