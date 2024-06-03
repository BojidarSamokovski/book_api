package main

import (
	"book_api/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
