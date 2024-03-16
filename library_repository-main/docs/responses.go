package docs

import (
	"github.com/geejjoo/library_repository/internal/infrastructure/response"
	"github.com/geejjoo/library_repository/internal/modules/library/controller"
)

// swagger:response commonResponse
type commonResponse struct {
	// in: body
	Body response.Response
}

// swagger:response authorCreateResponse
type authorCreateResponse struct {
	// in: body
	Body controller.CreateAuthorResponse
}

// swagger:response authorListResponse
type authorListResponse struct {
	// in: body
	Body controller.GetAllAuthorsResponse
}

// swagger:response bookCreateResponse
type bookCreateResponse struct {
	// in: body
	Body controller.CreateBookResponse
}

// swagger:response bookListResponse
type bookListResponse struct {
	// in: body
	Body controller.GetAllBooksResponse
}

// swagger:response userCreateResponse
type userCreateResponse struct {
	// in: body
	Body controller.CreateUserResponse
}

// swagger:response userListResponse
type userListResponse struct {
	// in: body
	Body controller.GetAllUsersResponse
}

// swagger:response ratingListResponse
type ratingListResponse struct {
	// in: body
	Body controller.GetRatingResponse
}

// swagger:response rentResponse
type rentResponse struct {
	// in: body
	Body controller.RentBookResponse
}

// swagger:response returnResponse
type returnResponse struct {
	// in: body
	Body controller.ReturnBookResponse
}
