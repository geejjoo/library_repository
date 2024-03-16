package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/geejjoo/library_repository/internal/models"
	"testing"
	"time"
)

func GenerateUser() models.UserDTO {
	return models.UserDTO{
		ID:          0,
		FirstName:   gofakeit.FirstName(),
		LastName:    gofakeit.LastName(),
		RentedBooks: nil,
	}
}

func TestUserCreate(t *testing.T) {
	for i := 0; i < 100; i++ {
		Do("POST", "http://localhost:8080/user", GenerateUser())
		time.Sleep(20 * time.Millisecond)
	}
}

func TestUserList(t *testing.T) {
	Do("GET", "http://localhost:8080/user", nil)
}
