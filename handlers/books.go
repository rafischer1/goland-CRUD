package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/rafischer1/goCRUD/models"
)

func GetAll(w http.ResponseWriter, req *http.Request) {
  
	data := models.GetAllBooks()
	vars := mux.Vars(req)
	
	//return the data
	fmt.Printf("d: %+v", data)
	w.WriteHeader(http.StatusOK)
  w.Write([]byte(vars["Content"]))
	fmt.Fprint(w, "Content:", vars["Content"], data)
}




//Rest of REST routes
func PostBook(w http.ResponseWriter, req *http.Request) {
  fmt.Println("In the handler post")
  
  data := models.PostBook()
  vars := mux.Vars(req)
  fmt.Fprint(w, "Content:", vars["Content"], data)
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
