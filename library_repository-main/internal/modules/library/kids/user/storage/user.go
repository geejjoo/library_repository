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

type UserStorage struct {
	adapter dao.Adapter
}

func NewUserStorage(sqlAdapter dao.Adapter) *UserStorage {
	return &UserStorage{
		adapter: sqlAdapter,
	}
}

func (u *UserStorage) Create(ctx context.Context, user models.UserDTO) error {
	err := u.adapter.Create(ctx, &user)
	if err != nil {
		if strings.Contains(err.Error(), db.AlreadyExists) {
			return fmt.Errorf(errors.StorageAlreadyExistsError)
		} else {
			return nil
		}
	}
	return nil
}

func (u *UserStorage) List(ctx context.Context) (out []models.UserDTO, err error) {
	var temp []struct {
		ID          int     `json:"id" db:"id" db_type:"SERIAL PRIMARY KEY"`
		FirstName   string  `json:"firstname" db:"firstname" db_type:"text"`
		LastName    string  `json:"lastname" db:"lastname" db_type:"text"`
		RentedBooks []uint8 `json:"rentedbooks" db:"rentedbooks" db_type:"text"`
	}
	err = u.adapter.List(ctx, &temp, &models.UserDTO{}, dao.Condition{})
	if err != nil {
		return nil, err
	}
	out = make([]models.UserDTO, len(temp))
	for i, v := range temp {
		var books []models.BookDTO
		json.Unmarshal(temp[i].RentedBooks, &books)
		out[i].SetID(v.ID)
		out[i].SetFirstName(v.FirstName)
		out[i].SetLastName(v.LastName)
		out[i].RentedBooks = books
	}
	return out, nil
}
func (u *UserStorage) Update(ctx context.Context, user models.UserDTO) error {
	err := u.adapter.Update(ctx, &user, dao.Condition{Equal: map[string]interface{}{
		"id": user.GetID(),
	}})
	if err != nil {
		return err
	}
	return nil
}

func (u *UserStorage) GetByID(ctx context.Context, id int) (out models.UserDTO, err error) {
	var temp []struct {
		ID          int     `json:"id" db:"id" db_type:"SERIAL PRIMARY KEY"`
		FirstName   string  `json:"firstname" db:"firstname" db_type:"text"`
		LastName    string  `json:"lastname" db:"lastname" db_type:"text"`
		RentedBooks []uint8 `json:"rentedbooks" db:"rentedbooks" db_type:"text"`
	}
	err = u.adapter.List(ctx, &temp, &models.UserDTO{}, dao.Condition{Equal: map[string]interface{}{
		"id": id,
	}})
	outs := make([]models.UserDTO, len(temp))
	for i, v := range temp {
		var books []models.BookDTO
		json.Unmarshal(temp[i].RentedBooks, &books)
		outs[i].SetID(v.ID)
		outs[i].SetFirstName(v.FirstName)
		outs[i].SetLastName(v.LastName)
		outs[i].RentedBooks = books
	}
	if err != nil || len(outs) == 0 {
		return models.UserDTO{}, err
	}
	return outs[0], nil
}
