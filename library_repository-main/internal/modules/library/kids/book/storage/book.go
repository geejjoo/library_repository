package storage

import (
	"context"
	"fmt"
	"github.com/geejjoo/library_repository/internal/db/adapter/dao"
	"github.com/geejjoo/library_repository/internal/infrastructure/db"
	"github.com/geejjoo/library_repository/internal/infrastructure/errors"
	"github.com/geejjoo/library_repository/internal/models"
	"strings"
)

type BookStorage struct {
	adapter dao.Adapter
}

func NewBookStorage(sqlAdapter dao.Adapter) *BookStorage {
	return &BookStorage{
		adapter: sqlAdapter,
	}
}

func (b *BookStorage) Create(ctx context.Context, book models.BookDTO) error {
	err := b.adapter.Create(ctx, &book)
	if err != nil {
		if strings.Contains(err.Error(), db.AlreadyExists) {
			return fmt.Errorf(errors.StorageAlreadyExistsError)
		} else {
			return err
		}
	}
	return nil
}

func (b *BookStorage) List(ctx context.Context) (out []models.BookDTO, err error) {
	err = b.adapter.List(ctx, &out, &models.BookDTO{}, dao.Condition{})
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (b *BookStorage) Update(ctx context.Context, book models.BookDTO) error {
	err := b.adapter.Update(ctx, &book, dao.Condition{Equal: map[string]interface{}{
		"id": book.GetID(),
	}})
	if err != nil {
		return err
	}
	return nil
}

func (b *BookStorage) GetByID(ctx context.Context, id int) (out models.BookDTO, err error) {
	var outs []models.BookDTO
	err = b.adapter.List(ctx, &outs, &models.BookDTO{}, dao.Condition{Equal: map[string]interface{}{
		"id": id,
	}})
	if err != nil || len(outs) == 0 {
		return models.BookDTO{}, err
	}
	return outs[0], nil
}
