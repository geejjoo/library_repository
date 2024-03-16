package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Uni(method, path string, body []byte) (*http.Response, error) {
	request, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	if err != nil {
		return &http.Response{}, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return &http.Response{}, err
	}
	return response, nil
}
func Do(method, url string, body any) {
	marshal, _ := json.Marshal(body)
	response, _ := Uni(method, url, marshal)
	all, _ := io.ReadAll(response.Body)
	fmt.Println(string(all))
}
