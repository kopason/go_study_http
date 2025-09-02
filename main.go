package main

import (
	"go_study_http/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	log.Println("Server starting at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
