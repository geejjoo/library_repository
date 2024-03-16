package router

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSwaggerUI(t *testing.T) {
	req, err := http.NewRequest("GET", "/swagger", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SwaggerUI)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "text/html; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, expectedContentType)
	}
}

func TestStaticHandler(t *testing.T) {
	//
	//Нужно запустить главный сервер
	//
	req, _ := http.NewRequest("GET", "http://localhost:8080/static/swagger.json?1705087308", nil)
	response, err := http.DefaultClient.Do(req)
	bytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bytes))
	if err != nil {
		return
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", response.StatusCode, http.StatusOK)
	}
}

func TestStaticHandlerNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/static/nonexistentfile.json", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StaticHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}
