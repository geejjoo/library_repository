package docs

import "github.com/geejjoo/library_repository/internal/modules/library/controller"

// swagger:route POST /user User userCreateRequest
// Добавление нового пользователя
// responses:
// 200: userCreateResponse
// 400: commonResponse
// 500: commonResponse

// swagger:parameters userCreateRequest
type userCreateRequest struct {
	//in:body
	//required: true
	Body controller.CreateUserRequest
}

// swagger:route GET /user User userListRequest
// Получение списка пользователей
// responses:
// 200: userListResponse
// 400: commonResponse
// 500: commonResponse

// swagger:parameters userListRequest
type userListRequest struct {
}
