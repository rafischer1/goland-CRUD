package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"errors"

	"github.com/rafischer1/goCRUD/models"
)

type Book struct {
	id    int64  `json:"id"`
	title string `json:"title"`
}

func GetAll(w http.ResponseWriter, req *http.Request) {

	data := models.GetAllBooks()
	vars := mux.Vars(req)

	//return the data
	fmt.Printf("d: %+v", data)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(vars["Content"]))
	fmt.Fprintf(w, "Content:%s", vars["Content"], data)
}

//Rest of REST routes
func PostBook(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("In the handler post rer", req)
	var book Book
	bookTitle := mux.Vars(req)["title"]
	data := models.GetAllBooks()
	
	if bookTitle == "" {
	errors.New("user id cannot be empty.")
}

	book.id, _ = strconv.ParseInt(req.FormValue("id"), 10, 64)
	book.title = req.FormValue("title")
	fmt.Println("req book handler:", book, data)
	// fmt.Fprint(w, "Content:", vars["Content"], data)
}

func EditBook(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In the handler edit")

	data := models.EditBook()
	vars := mux.Vars(req)
	fmt.Fprint(w, "Content:", vars["Content"], data)
}

func DeleteBook(w http.ResponseWriter, req *http.Request) {
	fmt.Println("In the handler delete")

	data := models.DeleteBook()
	vars := mux.Vars(req)
	fmt.Fprint(w, "Content:", vars["Content"], data)
}
