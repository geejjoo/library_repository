// Package docs GolibraryAPI.
//
// # Сервис библиотеки
//
// ### При авторизации указывайте префикс `"Bearer"`, пример: `"Bearer yourtoken12345"`
//
//	 Schemes: http, https
//	 Host: localhost:8080
//	 BasePath: /
//	 Version: 0.0.1
//	 Contact: yane.neya@mail.ru
//
//		Consumes:
//		- application/json
//		- multipart/form-data
//
//		Produces:
//		- application/json
//
//		Security:
//		- basic
//
//
//		SecurityDefinitions:
//		  Bearer:
//		    type: apiKey
//		    name: Authorization
//		    in: header
//
// swagger:meta
package docs

//go:generate swagger generate spec -o ../public/swagger.json --scan-models
