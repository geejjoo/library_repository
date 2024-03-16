package router

import (
	"github.com/geejjoo/library_repository/internal/infrastructure/component"
	"github.com/geejjoo/library_repository/internal/modules"
	"github.com/geejjoo/library_repository/internal/router"
	"github.com/go-chi/chi/v5"
)

func NewRouter(controllers *modules.Controllers, components *component.Components) *chi.Mux {
	r := chi.NewRouter()
	setDefaultRoutes(r)
	r.Mount("/", router.NewApiRouter(controllers, components.Guarder))
	return r
}

func setDefaultRoutes(r *chi.Mux) {
	r.Get("/swagger/index.html", SwaggerUI)
	r.Get("/static/*", StaticHandler)
}
