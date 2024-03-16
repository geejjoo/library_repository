package modules

import (
	"github.com/geejjoo/library_repository/internal/db/adapter/dao"
	astorage "github.com/geejjoo/library_repository/internal/modules/library/kids/author/storage"
	bstorage "github.com/geejjoo/library_repository/internal/modules/library/kids/book/storage"
	ustorage "github.com/geejjoo/library_repository/internal/modules/library/kids/user/storage"
)

type Storages struct {
	Author astorage.Authorer
	Book   bstorage.Booker
	User   ustorage.Userer
}

func NewStorages(sqlAdapter dao.Adapter) *Storages {
	return &Storages{
		Author: astorage.NewAuthorStorage(sqlAdapter),
		Book:   bstorage.NewBookStorage(sqlAdapter),
		User:   ustorage.NewUserStorage(sqlAdapter),
	}
}
