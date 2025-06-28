package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/JosTENill/go-book-api/store"
    "github.com/JosTENill/go-book-api/models"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    created := store.CreateBook(book)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(created)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    book, err := store.GetBook(id)
    if err != nil {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(book)
}

func ListBooks(w http.ResponseWriter, r *http.Request) {
    books := store.ListBooks()
    json.NewEncoder(w).Encode(books)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    var book models.Book
    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    updated, err := store.UpdateBook(id, book)
    if err != nil {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(updated)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    if err := store.DeleteBook(id); err != nil {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}


