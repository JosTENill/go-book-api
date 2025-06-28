package tests

import (
    "testing"

    "github.com/JosTENill/go-book-api/models"
    "github.com/JosTENill/go-book-api/store"
)

func TestCreateBook(t *testing.T) {
    b := models.Book{Title: "Test", Author: "Me", Pages: 123}
    created := store.CreateBook(b)

    if created.ID == 0 {
        t.Errorf("Expected ID > 0")
    }
    if created.Title != b.Title {
        t.Errorf("Title mismatch: got %q, want %q", created.Title, b.Title)
    }
}

func TestUpdateBook(t *testing.T) {
    b := models.Book{Title: "Old", Author: "Me", Pages: 50}
    created := store.CreateBook(b)

    updated := models.Book{Title: "New", Author: "You", Pages: 100}
    result, err := store.UpdateBook(created.ID, updated)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    if result.Title != "New" {
        t.Errorf("Failed to update title: got %q, want %q", result.Title, "New")
    }
}

func TestDeleteBook(t *testing.T) {
    b := models.Book{Title: "Temp", Author: "X", Pages: 10}
    created := store.CreateBook(b)

    err := store.DeleteBook(created.ID)
    if err != nil {
        t.Errorf("Delete failed: %v", err)
    }
    
    if _, err := store.GetBook(created.ID); err == nil {
        t.Errorf("Expected error for deleted book, but got none")
    }
}


