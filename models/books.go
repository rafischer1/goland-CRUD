package models

import (
	"database/sql"
	"log"
  "fmt"
)

type Book struct {
	id int64 `json:"id"`
	title string `json:"title"`
}


func GetAllBooks() []Book {

  connStr := "user=artiefischer dbname=golangtest sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
  defer db.Close()

	rows, err := db.Query("SELECT * FROM books")  
	
	defer rows.Close()
	var books []Book
	for rows.Next() {
		book := Book{}
		rows.Scan(&book.id, &book.title)
	  books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return books
}


//REST of rest Book routes
func PostBook() []Book {

  db, err := sql.Open("postgres", "user=artiefischer dbname=golangtest sslmode=disable")
  if err != nil {
    panic(err)
  }
var book []Book
	rows, err := db.Query(
	"INSERT INTO books (title) VALUES ($1)",
	book,
)
defer rows.Close()


// fmt.Println("In the model", book)
  // defer rows.Close()
  // var books []Book
  // for rows.Next() {
  //   book := Book{}
  //   rows.Scan(&book.id, &book.title)
  //   books = append(books, book)
  // }
  // if err := rows.Err(); err != nil {
  //   log.Fatal(err)
  // }

  return book
}


func EditBook() []Book {
  fmt.Println("In the model edit")
  db, err := sql.Open("postgres", "user=artiefischer dbname=golangtest sslmode=disable")
  if err != nil {
    panic(err)
  }
  
  rows, err := db.Query("SELECT * FROM books")
  
  defer rows.Close()
  var books []Book
  for rows.Next() {
    book := Book{}
    rows.Scan(&book.id, &book.title)
    books = append(books, book)
  }
  if err := rows.Err(); err != nil {
    log.Fatal(err)
  }

  return books
}

func DeleteBook() []Book {
  fmt.Println("In the model delete")
  db, err := sql.Open("postgres", "user=artiefischer dbname=golangtest sslmode=disable")
  if err != nil {
    panic(err)
  }
  
  rows, err := db.Query("SELECT * FROM books")
  
  defer rows.Close()
  var books []Book
  for rows.Next() {
    book := Book{}
    rows.Scan(&book.id, &book.title)
    books = append(books, book)
  }
  if err := rows.Err(); err != nil {
    log.Fatal(err)
  }

  return books
}
