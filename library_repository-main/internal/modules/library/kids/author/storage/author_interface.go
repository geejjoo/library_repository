package storage

import (
	"context"
	"github.com/geejjoo/library_repository/internal/models"
)

type Authorer interface {
	Create(ctx context.Context, author models.AuthorDTO) error
	List(ctx context.Context) (out []models.AuthorDTO, err error)
	Update(ctx context.Context, author models.AuthorDTO) error
	GetByID(ctx context.Context, id int) (out models.AuthorDTO, err error)
}
