package main

import (
	"github.com/AlbertMorenoDEV/go-ddd-playground/pkg/todo"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server...")

	router := todo.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
