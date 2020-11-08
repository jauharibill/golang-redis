package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestStore(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestIndex(t *testing.T) {
	req, rec := request(http.MethodGet, "/", nil)
	Index(rec, req)
	resp := rec.Result()

	assert.Equal(t, http.StatusConflict, resp.StatusCode)
}

func TestUpdate(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestDelete(t *testing.T) {
	assert.Equal(t, 4, 4)
}
