package integration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/wI2L/fizz"
)

// Get sends an HTTP request with the GET method transforms the JSON response
func Get[T any](handler *fizz.Fizz, result *T, url string) *httptest.ResponseRecorder {
	return request(handler, result, "GET", url, nil)
}

// Post sends an HTTP request with the POST method with a body and transforms the JSON response
func Post[T any](handler *fizz.Fizz, result *T, url string, body io.Reader) *httptest.ResponseRecorder {
	return request(handler, result, "POST", url, body)
}

// Put sends an HTTP request with the PUT method with a body and transforms the JSON response
func Put[T any](handler *fizz.Fizz, result *T, url string, body io.Reader) *httptest.ResponseRecorder {
	return request(handler, result, "PUT", url, body)
}

// Patch sends an HTTP request with the PATCH method with a body and transforms the JSON response
func Patch[T any](handler *fizz.Fizz, result *T, url string, body io.Reader) *httptest.ResponseRecorder {
	return request(handler, result, "PATCH", url, body)
}

// Delete sends an HTTP request with the DELETE method transforms the JSON response
func Delete[T any](handler *fizz.Fizz, result *T, url string) *httptest.ResponseRecorder {
	return request(handler, result, "DELETE", url, nil)
}

func request[T any](handler *fizz.Fizz, result *T, method string, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, fmt.Sprintf(path), body)
	res := httptest.NewRecorder()

	handler.Engine().ServeHTTP(res, req)

	if result != nil {
		if err := json.Unmarshal(res.Body.Bytes(), result); err != nil {
			panic(fmt.Errorf("Error deserializing JSON response: %v", err))
		}
	}

	return res
}
