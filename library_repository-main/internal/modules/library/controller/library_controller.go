package controller

import (
	"encoding/json"
	"fmt"
	"github.com/geejjoo/library_repository/internal/infrastructure/component"
	"github.com/geejjoo/library_repository/internal/infrastructure/errors"
	"github.com/geejjoo/library_repository/internal/infrastructure/responder"
	"github.com/geejjoo/library_repository/internal/models"
	"github.com/geejjoo/library_repository/internal/modules/library/service"
	"io"
	"net/http"
	"net/url"
)

type Librarian interface {
	CreateAuthor(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	CreateBook(w http.ResponseWriter, r *http.Request)
	RentBook(w http.ResponseWriter, r *http.Request)
	ReturnBook(w http.ResponseWriter, r *http.Request)
	GetRating(w http.ResponseWriter, r *http.Request)
	GetAllAuthors(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetAllBooks(w http.ResponseWriter, r *http.Request)
}

type Library struct {
	service   service.Librarian
	responder responder.Responder
}

func NewLibrary(service service.Librarian, components *component.Components) Librarian {
	return &Library{
		service:   service,
		responder: components.Responder,
	}
}

func (l *Library) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var createAuthorRequest CreateAuthorRequest
	err := l.JSONDecoder(r.Body, &createAuthorRequest)
	if err != nil {
		l.responder.ErrorBadRequest(w, err)
		return
	}
	err = l.service.CreateAuthor(r.Context(), models.AuthorDTO(createAuthorRequest))
	if err != nil {
		l.responder.ErrorInternal(w, err)
		return
	}
	l.responder.OutputJSON(w, CreateAuthorResponse{
		Success: true,
		Message: AuthorSuccessfullyCreated,
		Data:    nil,
	})
}

func (l *Library) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserRequest CreateUserRequest
	err := l.JSONDecoder(r.Body, &createUserRequest)
	if err != nil {
		l.responder.ErrorBadRequest(w, err)
		return
	}
	err = l.service.CreateUser(r.Context(), models.UserDTO(createUserRequest))
	if err != nil {
		l.responder.ErrorInternal(w, err)
		return
	}
	l.responder.OutputJSON(w, CreateUserResponse{
		Success: true,
		Message: UserSuccessfullyCreated,
		Data:    nil,
	})
}

func (l *Library) CreateBook(w http.ResponseWriter, r *http.Request) {
	var createBookRequest CreateBookRequest
	err := l.JSONDecoder(r.Body, &createBookRequest)
	if err != nil {
		l.responder.ErrorBadRequest(w, err)
		return
	}
	err = l.service.CreateBook(r.Context(), models.BookDTO(createBookRequest))
	if err != nil {
		l.responder.ErrorInternal(w, err)
		return
	}
	l.responder.OutputJSON(w, CreateBookResponse{
		Success: true,
		Message: BookSuccessfullyCreated,
		Data:    nil,
	})
}

func (l *Library) RentBook(w http.ResponseWriter, r *http.Request) {
	parser, err := url.Parse(r.URL.String())
	if err != nil {
		l.responder.ErrorBadRequest(w, fmt.Errorf(errors.HandlerParamsError))
		return
	}
	queryParser := parser.Query()
	bookID := queryParser.Get("bookID")
	userID := queryParser.Get("userID")
	if !l.ValidParams(bookID, userID) {
		l.responder.ErrorBadRequest(w, fmt.Errorf(errors.HandlerParamsError))
		return
	}
	err = l.service.RentBook(r.Context(), bookID, userID)
	if err != nil {
		l.responder.ErrorInternal(w, err)
		return
	}
	l.responder.OutputJSON(w, RentBookResponse{
		Success: true,
		Message: BookSuccessfullyRented,
		Data:    nil,
	})
}

func (l *Library) ReturnBook(w http.ResponseWriter, r *http.Request) {
	parser, err := url.Parse(r.URL.String())
	if err != nil {
		l.responder.ErrorBadRequest(w, fmt.Errorf(errors.HandlerParamsError))
		return
	}
	queryParser := parser.Query()
	bookID := queryParser.Get("bookID")
	userID := queryParser.Get("userID")
	if !l.ValidParams(bookID, userID) {
		l.responder.ErrorBadRequest(w, fmt.Errorf(errors.HandlerParamsError))
		return
	}
	err = l.service.ReturnBook(r.Context(), bookID, userID)
	if err != nil {
		l.responder.ErrorInternal(w, err)
		return
	}
	l.responder.OutputJSON(w, ReturnBookResponse{
		Success: true,
		Message: BookSuccessfullyReturned,
		Data:    nil,
	})
}

func (l *Library) GetRating(w http.ResponseWriter, r *http.Request) {
	rating, err := l.service.GetRating(r.Context())
	if err != nil {
		l.responder.ErrorInternal(w, err)
		return
	}
	l.responder.OutputJSON(w, GetRatingResponse{
		Success: true,
		Message: RatingSuccessfullyPulled,
		Data:    rating,
	})
}

func (l *Library) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := l.service.GetAllAuthors(r.Context())
	if err != nil {
		l.responder.ErrorInternal(w, err)
		return
	}
	l.responder.OutputJSON(w, GetAllAuthorsResponse{
		Success: true,
		Message: AuthorsSuccessfullyPulled,
		Data:    authors,
	})
}

func (l *Library) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := l.service.GetAllUsers(r.Context())
	if err != nil {
		l.responder.ErrorInternal(w, err)
		return
	}
	l.responder.OutputJSON(w, GetAllUsersResponse{
		Success: true,
		Message: UsersSuccessfullyPulled,
		Data:    users,
	})
}

func (l *Library) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := l.service.GetAllBooks(r.Context())
	if err != nil {
		l.responder.ErrorInternal(w, err)
		return
	}
	l.responder.OutputJSON(w, GetAllBooksResponse{
		Success: true,
		Message: BooksSuccessfullyPulled,
		Data:    books,
	})
}

func (l *Library) JSONDecoder(request io.Reader, dest any) error {
	err := json.NewDecoder(request).Decode(dest)
	if err != nil {
		return err
	}
	return nil
}

func (l *Library) ValidParams(params ...string) bool {
	for _, p := range params {
		if p == "" {
			return false
		}
	}
	return true
}
