package docs

// swagger:route GET /rating Library ratingListRequest
// Получение рейтинга авторов
// responses:
// 200: ratingListResponse
// 400: commonResponse
// 500: commonResponse

// swagger:parameters ratingListRequest
type ratingListRequest struct {
}

// swagger:route GET /rent Library rentRequest
// Выдача книги
// responses:
// 200: rentResponse
// 400: commonResponse
// 500: commonResponse

// swagger:parameters rentRequest
type rentRequest struct {
	//in:query
	//required: true
	BookID string
	//in:query
	//required: true
	UserID string
}

// swagger:route GET /return Library returnRequest
// Возврат книги
// responses:
// 200: returnResponse
// 400: commonResponse
// 500: commonResponse

// swagger:parameters returnRequest
type returnRequest struct {
	//in:query
	//required: true
	BookID string
	//in:query
	//required: true
	UserID string
}
