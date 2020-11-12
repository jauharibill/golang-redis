package utilities

import (
	_ "github.com/stretchr/testify/assert"
	"github.com/vmihailenco/treemux"
	"io"
	_ "net/http"
	"net/http/httptest"
	"strconv"
)

func Request(method string, url string, payload io.Reader) (treemux.Request, *httptest.ResponseRecorder) {
	var request treemux.Request

	req := httptest.NewRequest(method, url, payload)
	req.Header.Set("Content-type", "application/json")
	rec := httptest.NewRecorder()

	request.Request = req

	return request, rec
}

func StrToInt(param string) int {
	result, _ := strconv.Atoi(param)

	return result
}

func IntToStr(param int) string {
	return strconv.Itoa(param)
}
