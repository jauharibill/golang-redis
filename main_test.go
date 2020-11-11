package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestStore(t *testing.T) {
	var user User
	controller := InitController(Conn())

	user.ID = 123
	user.Name = "Tanthowi"
	user.Age = 26

	out, _ := json.Marshal(user)
	payload := bytes.NewBuffer(out)

	req, rec := request(http.MethodPost, "/store", payload)

	controller.Store(rec, req)
	resp := rec.Result()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestShow(t *testing.T) {
	controller := InitController(Conn())
	req, rec := request(http.MethodGet, "/show?id=123", nil)

	controller.Show(rec, req)
	resp := rec.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUpdate(t *testing.T) {

	var user User
	controller := InitController(Conn())

	user.ID = 123
	user.Name = "Jauhari"
	user.Age = 26

	out, _ := json.Marshal(user)
	payload := bytes.NewBuffer(out)

	req, rec := request(http.MethodPost, "/update?id=123", payload)

	controller.Update(rec, req)
	resp := rec.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDelete(t *testing.T) {
	controller := InitController(Conn())

	req, rec := request(http.MethodDelete, "/delete?id=123", nil)

	controller.Delete(rec, req)

	resp := rec.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
