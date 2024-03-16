package service

import (
	"context"
	"github.com/geejjoo/library_repository/internal/models"
)

type Booker interface {
	CreateBook(ctx context.Context, book models.BookDTO) error
	GetAllBooks(ctx context.Context) (out []models.BookDTO, err error)
	UpdateBook(ctx context.Context, book models.BookDTO) error
	GetBookByID(ctx context.Context, id int) (out models.BookDTO, err error)
}
