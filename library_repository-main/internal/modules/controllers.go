package modules

import (
	"github.com/geejjoo/library_repository/internal/infrastructure/component"
	lcontroller "github.com/geejjoo/library_repository/internal/modules/library/controller"
)

type Controllers struct {
	Library lcontroller.Librarian
}

func NewControllers(services *Services, components *component.Components) *Controllers {
	return &Controllers{
		Library: lcontroller.NewLibrary(services.Library, components),
	}
}
