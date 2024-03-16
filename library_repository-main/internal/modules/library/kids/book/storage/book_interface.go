package storage

import (
	"context"
	"github.com/geejjoo/library_repository/internal/models"
)

type Booker interface {
	Create(ctx context.Context, book models.BookDTO) error
	List(ctx context.Context) (out []models.BookDTO, err error)
	Update(ctx context.Context, book models.BookDTO) error
	GetByID(ctx context.Context, id int) (out models.BookDTO, err error)
}
