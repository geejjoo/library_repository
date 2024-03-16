package service

import (
	"context"
	"github.com/geejjoo/library_repository/internal/models"
)

type Authorer interface {
	CreateAuthor(ctx context.Context, author models.AuthorDTO) error
	GetAllAuthors(ctx context.Context) (out []models.AuthorDTO, err error)
	UpdateAuthor(ctx context.Context, author models.AuthorDTO) error
	GetAuthorByID(ctx context.Context, id int) (out models.AuthorDTO, err error)
}
