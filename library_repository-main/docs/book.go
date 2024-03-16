package docs

import "github.com/geejjoo/library_repository/internal/modules/library/controller"

// swagger:route POST /book Book bookCreateRequest
// Добавление новой книги
// responses:
// 200: bookCreateResponse
// 400: commonResponse
// 500: commonResponse

// swagger:parameters bookCreateRequest
type bookCreateRequest struct {
	//in:body
	//required: true
	Body controller.CreateBookRequest
}

// swagger:route GET /book Book bookListRequest
// Получение списка книг
// responses:
// 200: bookListResponse
// 400: commonResponse
// 500: commonResponse

// swagger:parameters bookListRequest
type bookListRequest struct {
}
