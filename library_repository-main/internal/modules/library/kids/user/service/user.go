package service

import (
	"context"
	"fmt"
	"github.com/geejjoo/library_repository/internal/infrastructure/errors"
	"github.com/geejjoo/library_repository/internal/models"
	"github.com/geejjoo/library_repository/internal/modules/library/kids/user/storage"
	"go.uber.org/zap"
	"reflect"
)

type UserService struct {
	storage storage.Userer
	logger  *zap.Logger
}

func NewUserService(storage storage.Userer, logger *zap.Logger) *UserService {
	return &UserService{
		storage: storage,
		logger:  logger,
	}
}

func (u *UserService) CreateUser(ctx context.Context, user models.UserDTO) error {
	err := u.storage.Create(ctx, user)
	if err != nil {
		u.logger.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (u *UserService) GetAllUsers(ctx context.Context) (out []models.UserDTO, err error) {
	list, err := u.storage.List(ctx)
	if err != nil {
		u.logger.Error("", zap.Error(err))
		return nil, err
	}
	return list, nil
}

func (u *UserService) UpdateUser(ctx context.Context, user models.UserDTO) error {
	err := u.storage.Update(ctx, user)
	if err != nil {
		u.logger.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (u *UserService) GetUserByID(ctx context.Context, id int) (out models.UserDTO, err error) {
	byID, err := u.storage.GetByID(ctx, id)
	if err != nil || u.IsEmpty(byID, models.UserDTO{}) {
		u.logger.Error(errors.ServiceGetByIDError, zap.Error(err))
		return models.UserDTO{}, fmt.Errorf(errors.ServiceGetByIDError)
	}
	return byID, nil
}

func (u *UserService) IsEmpty(x, y any) bool {
	return reflect.DeepEqual(x, y)
}
