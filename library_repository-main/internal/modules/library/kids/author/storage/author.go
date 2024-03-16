package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/geejjoo/library_repository/internal/db/adapter/dao"
	"github.com/geejjoo/library_repository/internal/infrastructure/db"
	"github.com/geejjoo/library_repository/internal/infrastructure/errors"
	"github.com/geejjoo/library_repository/internal/models"
	"strings"
)

type AuthorStorage struct {
	adapter dao.Adapter
}

func NewAuthorStorage(sqlAdapter dao.Adapter) *AuthorStorage {
	return &AuthorStorage{
		adapter: sqlAdapter,
	}
}

func (a *AuthorStorage) Create(ctx context.Context, author models.AuthorDTO) error {
	err := a.adapter.Create(ctx, &author)
	if err != nil {
		if strings.Contains(err.Error(), db.AlreadyExists) {
			return fmt.Errorf(errors.StorageAlreadyExistsError)
		} else {
			return err
		}
	}
	return nil
}

func (a *AuthorStorage) List(ctx context.Context) (out []models.AuthorDTO, err error) {
	var temp []struct {
		ID        int     `json:"id" db:"id" db_type:"SERIAL PRIMARY KEY"`
		FirstName string  `json:"firstname" db:"firstname" db_type:"text"`
		LastName  string  `json:"lastname" db:"lastname" db_type:"text"`
		Books     []uint8 `json:"books" db:"books" db_type:"json"`
	}

	err = a.adapter.List(ctx, &temp, &models.AuthorDTO{}, dao.Condition{})
	if err != nil {
		return nil, err
	}
	out = make([]models.AuthorDTO, len(temp))
	for i, v := range temp {
		var books []models.BookDTO
		json.Unmarshal(temp[i].Books, &books)
		out[i].SetID(v.ID)
		out[i].SetFirstName(v.FirstName)
		out[i].SetLastName(v.LastName)
		out[i].Books = books
	}
	return out, nil
}

func (a *AuthorStorage) Update(ctx context.Context, author models.AuthorDTO) error {
	err := a.adapter.Update(ctx, &author, dao.Condition{Equal: map[string]interface{}{
		"id": author.GetID(),
	}})
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthorStorage) GetByID(ctx context.Context, id int) (out models.AuthorDTO, err error) {
	var temp []struct {
		ID        int     `json:"id" db:"id" db_type:"SERIAL PRIMARY KEY"`
		FirstName string  `json:"firstname" db:"firstname" db_type:"text"`
		LastName  string  `json:"lastname" db:"lastname" db_type:"text"`
		Books     []uint8 `json:"books" db:"books" db_type:"json"`
	}
	err = a.adapter.List(ctx, &temp, &models.AuthorDTO{}, dao.Condition{Equal: map[string]interface{}{
		"id": id,
	}})
	outs := make([]models.AuthorDTO, len(temp))
	for i, v := range temp {
		var books []models.BookDTO
		json.Unmarshal(temp[i].Books, &books)
		outs[i].SetID(v.ID)
		outs[i].SetFirstName(v.FirstName)
		outs[i].SetLastName(v.LastName)
		outs[i].Books = books
	}
	if err != nil || len(outs) == 0 {
		return models.AuthorDTO{}, err
	}
	return outs[0], nil
}
