package routes

import (
	"github.com/ashwinpnr/golang-samples/go-mysql-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(r *mux.Router) {
	r.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	r.HandleFunc("/book/{id}", controllers.GetBookByID).Methods("GET")
	r.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")
}
