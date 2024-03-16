package router

import (
	"github.com/geejjoo/library_repository/internal/infrastructure/middleware/auth/guarder"
	"github.com/geejjoo/library_repository/internal/modules"
	"github.com/go-chi/chi"
	"net/http"
)

func NewApiRouter(controllers *modules.Controllers, protector guarder.Guarder) http.Handler {
	r := chi.NewRouter()
	lcontroller := controllers.Library
	r.Route("/", func(r chi.Router) {
		r.Post("/author", lcontroller.CreateAuthor)
		r.Post("/book", lcontroller.CreateBook)
		r.Post("/user", lcontroller.CreateUser)
		r.Get("/rent", lcontroller.RentBook)
		r.Get("/return", lcontroller.ReturnBook)
		r.Get("/rating", lcontroller.GetRating)
		r.Get("/author", lcontroller.GetAllAuthors)
		r.Get("/book", lcontroller.GetAllBooks)
		r.Get("/user", lcontroller.GetAllUsers)
	})
	return r
}
