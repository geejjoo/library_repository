package controller

import (
	"github.com/geejjoo/library_repository/internal/infrastructure/response"
	"github.com/geejjoo/library_repository/internal/models"
)

const (
	AuthorSuccessfullyCreated = "author successfully created"
	UserSuccessfullyCreated   = "user successfully created"
	BookSuccessfullyCreated   = "book successfully created"

	BookSuccessfullyRented   = "book successfully rented"
	BookSuccessfullyReturned = "book successfully returned"

	RatingSuccessfullyPulled  = "rating successfully pulled"
	AuthorsSuccessfullyPulled = "authors successfully pulled"
	UsersSuccessfullyPulled   = "users successfully pulled"
	BooksSuccessfullyPulled   = "books successfully pulled"
)

type CreateAuthorRequest models.AuthorDTO
type CreateUserRequest models.UserDTO
type CreateBookRequest models.BookDTO

type CreateAuthorResponse response.Response
type CreateUserResponse response.Response
type CreateBookResponse response.Response

type RentBookResponse response.Response
type ReturnBookResponse response.Response

type GetRatingResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Data    []models.RatingDTO `json:"data"`
}
type GetAllAuthorsResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Data    []models.AuthorDTO `json:"data"`
}
type GetAllUsersResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Data    []models.UserDTO `json:"data"`
}
type GetAllBooksResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Data    []models.BookDTO `json:"data"`
}
