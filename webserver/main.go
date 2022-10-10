package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "hello from server")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

}

func main() {

	fileHandler := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	err := http.ListenAndServe("localhost:8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
