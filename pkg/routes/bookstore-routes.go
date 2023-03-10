package routes

import (
	"github.com/gorilla/mux"
	"github.com/KleitonBarone/Go-Bookstore/pkg/controllers"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId:[0-9]+}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId:[0-9]+}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId:[0-9]+}", controllers.DeleteBook).Methods("DELETE")
}
