package main

import (
	_ "github.com/stretchr/testify/assert"
	"io"
	"net/http"
	_ "net/http"
	"net/http/httptest"
	"strconv"
)

func request(method string, url string, payload io.Reader) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, payload)
	req.Header.Set("Content-type", "application/json")
	rec := httptest.NewRecorder()

	return req, rec
}

func StrToInt(param string) int {
	result, _ := strconv.Atoi(param)

	return result
}
