package main

import (
	"log"
	"net/http"

	"github.com/ashwinpnr/golang-samples/go-mysql-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Printf("Starting Server")
	if err := http.ListenAndServe(":9010", r); err != nil {
		log.Fatal(err)
	}

}
