package service

import (
	"context"
	"github.com/geejjoo/library_repository/internal/models"
)

type Librarian interface {
	CreateAuthor(ctx context.Context, author models.AuthorDTO) error
	CreateUser(ctx context.Context, user models.UserDTO) error
	CreateBook(ctx context.Context, book models.BookDTO) error
	RentBook(ctx context.Context, bookID, userID string) error
	ReturnBook(ctx context.Context, bookID, userID string) error
	GetRating(ctx context.Context) (out []models.RatingDTO, err error)
	GetAllAuthors(ctx context.Context) (out []models.AuthorDTO, err error)
	GetAllUsers(ctx context.Context) (out []models.UserDTO, err error)
	GetAllBooks(ctx context.Context) (out []models.BookDTO, err error)
}
