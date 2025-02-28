package router

import (
	"book_api/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
	return router
}
