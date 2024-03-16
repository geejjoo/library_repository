package service

import (
	"context"
	"github.com/geejjoo/library_repository/internal/models"
)

type Userer interface {
	CreateUser(ctx context.Context, user models.UserDTO) error
	GetAllUsers(ctx context.Context) (out []models.UserDTO, err error)
	UpdateUser(ctx context.Context, user models.UserDTO) error
	GetUserByID(ctx context.Context, id int) (out models.UserDTO, err error)
}
