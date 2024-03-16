package docs

import "github.com/geejjoo/library_repository/internal/modules/library/controller"

// swagger:route POST /author Author authorCreateRequest
// Добавление нового автора
// responses:
// 200: authorCreateResponse
// 400: commonResponse
// 500: commonResponse

// swagger:parameters authorCreateRequest
type authorCreateRequest struct {
	//in:body
	//required: true
	Body controller.CreateAuthorRequest
}

// swagger:route GET /author Author authorListRequest
// Получение списка авторов
// responses:
// 200: authorListResponse
// 400: commonResponse
// 500: commonResponse

// swagger:parameters authorListRequest
type authorListRequest struct {
}
