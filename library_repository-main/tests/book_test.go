package main

import (
	"github.com/geejjoo/library_repository/internal/models"
	"testing"
)

func TestBookCreate(t *testing.T) {
	Do("POST", "http://localhost:8080/book", models.BookDTO{
		ID:       0,
		Title:    "MyBook",
		AuthorID: 1,
		ReaderID: 0,
	})
}
func TestBookList(t *testing.T) {
	Do("GET", "http://localhost:8080/book", nil)
}
