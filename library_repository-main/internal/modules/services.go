package modules

import (
	"github.com/geejjoo/library_repository/internal/infrastructure/component"
	aservice "github.com/geejjoo/library_repository/internal/modules/library/kids/author/service"
	bservice "github.com/geejjoo/library_repository/internal/modules/library/kids/book/service"
	uservice "github.com/geejjoo/library_repository/internal/modules/library/kids/user/service"
	lservice "github.com/geejjoo/library_repository/internal/modules/library/service"
)

type Services struct {
	Library lservice.Librarian
}

func NewServices(storages *Storages, components *component.Components) *Services {
	author := aservice.NewAuthorService(storages.Author, components.Logger)
	book := bservice.NewBookService(storages.Book, components.Logger)
	user := uservice.NewUserService(storages.User, components.Logger)
	return &Services{
		Library: lservice.NewLibraryService(author, book, user, components.Logger),
	}
}
