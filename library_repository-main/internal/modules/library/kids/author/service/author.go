package service

import (
	"context"
	"fmt"
	"github.com/geejjoo/library_repository/internal/infrastructure/errors"
	"github.com/geejjoo/library_repository/internal/models"
	"github.com/geejjoo/library_repository/internal/modules/library/kids/author/storage"
	"go.uber.org/zap"
	"reflect"
)

type AuthorService struct {
	storage storage.Authorer
	logger  *zap.Logger
}

func NewAuthorService(storage storage.Authorer, logger *zap.Logger) *AuthorService {
	return &AuthorService{
		storage: storage,
		logger:  logger,
	}
}

func (a *AuthorService) CreateAuthor(ctx context.Context, author models.AuthorDTO) error {
	err := a.storage.Create(ctx, author)
	if err != nil {
		a.logger.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (a *AuthorService) GetAllAuthors(ctx context.Context) (out []models.AuthorDTO, err error) {
	out, err = a.storage.List(ctx)
	if err != nil {
		a.logger.Error("", zap.Error(err))
		return nil, err
	}
	return out, nil
}

func (a *AuthorService) UpdateAuthor(ctx context.Context, author models.AuthorDTO) error {
	err := a.storage.Update(ctx, author)
	if err != nil {
		a.logger.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (a *AuthorService) GetAuthorByID(ctx context.Context, id int) (out models.AuthorDTO, err error) {
	byID, err := a.storage.GetByID(ctx, id)
	if err != nil || a.IsEmpty(byID, models.AuthorDTO{}) {
		a.logger.Error(errors.ServiceGetByIDError, zap.Error(err))
		return models.AuthorDTO{}, fmt.Errorf(errors.ServiceGetByIDError)
	}
	return byID, nil
}

func (a *AuthorService) IsEmpty(x, y any) bool {
	return reflect.DeepEqual(x, y)
}
