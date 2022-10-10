package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	gorillamux "github.com/gorilla/mux"
)

type Movie struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
}

var movies []Movie

func defaultMovies(movies []Movie) []Movie {
	movies = append(movies, Movie{
		Id:    "1",
		Title: "Deafult Movie One",
		Director: &Director{
			Firstname: "William",
		},
	})
	movies = append(movies, Movie{
		Id:    "2",
		Title: "Deafult Movie Two",
		Director: &Director{
			Firstname: "James",
		},
	})
	return movies
}

func getMovies(w http.ResponseWriter,
	r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func RemoveMovieIndex(s []Movie, index int) []Movie {
	return append(s[:index], s[index+1:]...)
}

func deleteMovie(w http.ResponseWriter,
	r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			RemoveMovieIndex(movies, index)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter,
	r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter,
	r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, "Couldnt decode movie entry for create", http.StatusNotAcceptable)
	}
	movie.Id = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter,
	r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			var movie Movie
			RemoveMovieIndex(movies, index)
			err := json.NewDecoder(r.Body).Decode(&movie)
			if err != nil {
				http.Error(w, "Couldnt decode movie entry for update", http.StatusNotAcceptable)
			}
			movie.Id = item.Id
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := gorillamux.NewRouter()
	movies = defaultMovies(movies)
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movie", createMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")
	log.Printf("Starting Server")

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}

}
