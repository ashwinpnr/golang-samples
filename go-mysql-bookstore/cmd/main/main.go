package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ashwinpnr/golang-samples/go-mysql-bookstore/pkg/config"
	"github.com/ashwinpnr/golang-samples/go-mysql-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Printf("Starting Server")
	serveraddr := ":" + strconv.Itoa(config.GetConfig().Server.Port)
	if err := http.ListenAndServe(serveraddr, r); err != nil {
		log.Fatal(err)
	}
}
