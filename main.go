package main

/*
*pq - A pure Go postgres driver for Go's database/sql package
*Mux - A powerful HTTP router and URL matcher for building Go web servers with
*strconv - A Go Core module for String conversion to INT
*encoding/json - A Core GO Module for decoding and encoding JSON
*net/http -Core GO Module  This package allows you to build HTTP servers in Go with its powerful compositional constructs
*gotenv - Load environment variables from .env or io.Reader in Go.
*database/sql - Provides generic Interface for database/Sql Interaction
*os - Package os provides a platform-independent interface to operating system functionality
 */

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/subosito/gotenv"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func init() {
	gotenv.Load()
}

func logFal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := mux.NewRouter()

	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}

func getBooks(w http.ResponseWriter, r *http.Request) {

}

func getBook(w http.ResponseWriter, r *http.Request) {

}

func addBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func removeBook(w http.ResponseWriter, r *http.Request) {

}
