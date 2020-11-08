package main

import (
	"io"
	_ "net/http"
	"net/http/httptest"
)

func request(method string, url string, payload io.Reader) {
	req := httptest.NewRequest(method, url, payload)
	req.Header.Set("Content-type", "application/json")
}
