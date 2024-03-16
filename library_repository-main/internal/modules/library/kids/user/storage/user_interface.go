package storage

import (
	"context"
	"github.com/geejjoo/library_repository/internal/models"
)

type Userer interface {
	Create(ctx context.Context, user models.UserDTO) error
	List(ctx context.Context) (out []models.UserDTO, err error)
	Update(ctx context.Context, user models.UserDTO) error
	GetByID(ctx context.Context, id int) (out models.UserDTO, err error)
}
