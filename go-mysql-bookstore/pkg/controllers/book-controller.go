package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ashwinpnr/golang-samples/go-mysql-bookstore/pkg/models"
	"github.com/ashwinpnr/golang-samples/go-mysql-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

func GetBooks(w http.ResponseWriter,
	r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books := models.GetAllBooks()
	result, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "Get Books Json Marshall Error", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func CreateBook(w http.ResponseWriter,
	r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	book := models.Book{}
	utils.ParseBody(r, &book)
	createdBook := models.CreateBook(book)
	result, err := json.Marshal(createdBook)
	if err != nil {
		http.Error(w, "Create Book Json Marshall Error", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetBookByID(w http.ResponseWriter,
	r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bookId := params["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Get Book by ID Error converting book id to int64", http.StatusInternalServerError)
	}
	book := models.GetBookByID(ID)
	result, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Get Book by ID Json Marshall Error", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteBook(w http.ResponseWriter,
	r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bookId := params["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Delete Book by ID Error converting book id to int64", http.StatusInternalServerError)
	}
	book := models.DeleteBook(ID)
	result, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Delete Book by ID Json Marshall Error", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateBook(w http.ResponseWriter,
	r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bookId := params["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Updated Book Error converting book id to int64", http.StatusInternalServerError)
	}
	booktoupdate := models.Book{}
	err = utils.ParseBody(r, &booktoupdate)
	if err != nil {
		http.Error(w, "Create Book Json Marshall Error", http.StatusInternalServerError)
	}
	updatedbook := models.UpdateBook(ID, booktoupdate)
	result, err := json.Marshal(updatedbook)
	if err != nil {
		http.Error(w, "Update Book Json Marshall Error", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
