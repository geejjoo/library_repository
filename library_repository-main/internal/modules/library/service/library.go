package service

import (
	"context"
	"fmt"
	"github.com/geejjoo/library_repository/internal/infrastructure/errors"
	"github.com/geejjoo/library_repository/internal/models"
	aservice "github.com/geejjoo/library_repository/internal/modules/library/kids/author/service"
	bservice "github.com/geejjoo/library_repository/internal/modules/library/kids/book/service"
	uservice "github.com/geejjoo/library_repository/internal/modules/library/kids/user/service"
	"go.uber.org/zap"
	"sort"
	"strconv"
)

type LibraryService struct {
	Author aservice.Authorer
	Book   bservice.Booker
	User   uservice.Userer
	Cache  map[int]int
	logger *zap.Logger
}

func NewLibraryService(author aservice.Authorer, book bservice.Booker, user uservice.Userer, logger *zap.Logger) *LibraryService {
	return &LibraryService{
		Author: author,
		Book:   book,
		User:   user,
		Cache:  map[int]int{},
		logger: logger,
	}
}

func (l *LibraryService) CreateAuthor(ctx context.Context, author models.AuthorDTO) error {
	err := l.Author.CreateAuthor(ctx, author)
	if err != nil {
		return err
	}
	return nil
}

func (l *LibraryService) CreateUser(ctx context.Context, user models.UserDTO) error {
	err := l.User.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (l *LibraryService) CreateBook(ctx context.Context, book models.BookDTO) error {
	author, err := l.Author.GetAuthorByID(ctx, book.GetAuthorID())
	if err != nil {
		l.logger.Error("No such author", zap.Error(err))
		return fmt.Errorf("no such author")
	}
	err = l.Book.CreateBook(ctx, book)
	if err != nil {
		return err
	}
	author.SetBooks(book)
	err = l.Author.UpdateAuthor(ctx, author)
	if err != nil {
		return err
	}
	return nil
}

func (l *LibraryService) RentBook(ctx context.Context, bookID, userID string) error {
	tempMap := l.ParseIDS(bookID, userID)
	intBookID := tempMap[bookID]
	intUserID := tempMap[userID]
	if intBookID == -1 || intUserID == -1 {
		l.logger.Error(errors.ServiceIDEncodeError)
		return fmt.Errorf(errors.ServiceIDEncodeError)
	}
	getBook, err := l.Book.GetBookByID(ctx, intBookID)
	if err != nil {
		l.logger.Error("No such book", zap.Error(err))
		return fmt.Errorf("no such book")
	}
	getUser, err := l.User.GetUserByID(ctx, intUserID)
	if err != nil {
		l.logger.Error("No such user", zap.Error(err))
		return fmt.Errorf("no such user")
	}
	if getBook.ReaderID != 0 {
		l.logger.Error("The book is already taken")
		return fmt.Errorf("the book is already taken")
	}
	getBook.SetReaderID(getUser.GetID())
	err = l.Book.UpdateBook(ctx, getBook)
	if err != nil {
		return err
	}
	getUser.AddRentedBooks(getBook)
	err = l.User.UpdateUser(ctx, getUser)
	if err != nil {
		return err
	}
	l.Cache[getBook.AuthorID]++
	return nil
}

func (l *LibraryService) ReturnBook(ctx context.Context, bookID, userID string) error {
	tempMap := l.ParseIDS(bookID, userID)
	intBookID := tempMap[bookID]
	intUserID := tempMap[userID]
	if intBookID == -1 || intUserID == -1 {
		l.logger.Error(errors.ServiceIDEncodeError)
		return fmt.Errorf(errors.ServiceIDEncodeError)
	}
	getBook, err := l.Book.GetBookByID(ctx, intBookID)
	if err != nil {
		l.logger.Error("No such book", zap.Error(err))
		return fmt.Errorf("no such book")
	}
	getUser, err := l.User.GetUserByID(ctx, intUserID)
	if err != nil {
		l.logger.Error("No such user", zap.Error(err))
		return fmt.Errorf("no such user")
	}
	if getBook.GetReaderID() != getUser.GetID() {
		l.logger.Error("Book and user mismatch")
		return fmt.Errorf("book and user mismatch")
	}
	getBook.SetReaderID(0)
	err = l.Book.UpdateBook(ctx, getBook)
	if err != nil {
		return err
	}
	getUser.DeleteRentedBooks(getBook.GetID())
	err = l.User.UpdateUser(ctx, getUser)
	if err != nil {
		return err
	}
	return nil
}

func (l *LibraryService) GetRating(ctx context.Context) (out []models.RatingDTO, err error) {
	authors, err := l.GetAllAuthors(ctx)
	if err != nil {
		return nil, err
	}
	for _, a := range authors {
		out = append(out, models.RatingDTO{
			Count:  l.Cache[a.GetID()],
			Author: a,
		})
	}
	if len(out) > 1 {
		sort.Slice(out, func(i, j int) bool {
			return out[i].Count > out[j].Count
		})
	}
	return out, nil
}

func (l *LibraryService) GetAllAuthors(ctx context.Context) (out []models.AuthorDTO, err error) {
	out, err = l.Author.GetAllAuthors(ctx)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (l *LibraryService) GetAllUsers(ctx context.Context) (out []models.UserDTO, err error) {
	out, err = l.User.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (l *LibraryService) GetAllBooks(ctx context.Context) (out []models.BookDTO, err error) {
	out, err = l.Book.GetAllBooks(ctx)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (l *LibraryService) ParseIDS(ids ...string) (out map[string]int) {
	out = make(map[string]int)
	for _, value := range ids {
		atoi, err := strconv.Atoi(value)
		if err != nil {
			out[value] = -1
			continue
		}
		out[value] = atoi
	}
	return out
}
