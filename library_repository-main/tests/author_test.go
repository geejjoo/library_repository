package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/geejjoo/library_repository/internal/models"
	"testing"
	"time"
)

func GenerateAuthor() models.AuthorDTO {
	return models.AuthorDTO{
		ID:        0,
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Books:     nil,
	}
}

func TestAuthorCreate(t *testing.T) {
	for i := 0; i < 100; i++ {
		Do("POST", "http://localhost:8080/author", GenerateAuthor())
		time.Sleep(20 * time.Millisecond)
	}
}
func TestAuthorList(t *testing.T) {
	Do("GET", "http://localhost:8080/author", nil)
}
