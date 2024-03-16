package service

import (
	"context"
	"fmt"
	"github.com/geejjoo/library_repository/internal/infrastructure/errors"
	"github.com/geejjoo/library_repository/internal/models"
	"github.com/geejjoo/library_repository/internal/modules/library/kids/book/storage"
	"go.uber.org/zap"
	"reflect"
)

type BookService struct {
	storage storage.Booker
	logger  *zap.Logger
}

func NewBookService(storage storage.Booker, logger *zap.Logger) *BookService {
	return &BookService{
		storage: storage,
		logger:  logger,
	}
}

func (b *BookService) CreateBook(ctx context.Context, book models.BookDTO) error {
	err := b.storage.Create(ctx, book)
	if err != nil {
		b.logger.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (b *BookService) GetAllBooks(ctx context.Context) (out []models.BookDTO, err error) {
	list, err := b.storage.List(ctx)
	if err != nil {
		b.logger.Error("", zap.Error(err))
		return nil, err
	}
	return list, nil
}

func (b *BookService) UpdateBook(ctx context.Context, book models.BookDTO) error {
	err := b.storage.Update(ctx, book)
	if err != nil {
		b.logger.Error("", zap.Error(err))
		return err
	}
	return nil
}

func (b *BookService) GetBookByID(ctx context.Context, id int) (out models.BookDTO, err error) {
	byID, err := b.storage.GetByID(ctx, id)
	if err != nil || b.IsEmpty(byID, models.BookDTO{}) {
		b.logger.Error(errors.ServiceGetByIDError, zap.Error(err))
		return models.BookDTO{}, fmt.Errorf(errors.ServiceGetByIDError)
	}
	return byID, nil
}

func (b *BookService) IsEmpty(x, y any) bool {
	return reflect.DeepEqual(x, y)
}
