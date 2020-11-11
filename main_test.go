package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"golang-redis/controllers"
	"golang-redis/utilities"
	"net/http"
	"testing"
)

func TestStore(t *testing.T) {
	var user User
	controller := controllers.InitController(Conn())

	user.ID = 123
	user.Name = "Tanthowi"
	user.Age = 26

	out, _ := json.Marshal(user)
	payload := bytes.NewBuffer(out)

	req, rec := utilities.request(http.MethodPost, "/store", payload)

	controller.Store(rec, req)
	resp := rec.Result()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

func TestShow(t *testing.T) {
	controller := controllers.InitController(Conn())
	req, rec := utilities.request(http.MethodGet, "/show?id=123", nil)

	controller.Show(rec, req)
	resp := rec.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUpdate(t *testing.T) {

	var user User
	controller := controllers.InitController(Conn())

	user.ID = 123
	user.Name = "Jauhari"
	user.Age = 26

	out, _ := json.Marshal(user)
	payload := bytes.NewBuffer(out)

	req, rec := utilities.request(http.MethodPost, "/update?id=123", payload)

	controller.Update(rec, req)
	resp := rec.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDelete(t *testing.T) {
	controller := controllers.InitController(Conn())

	req, rec := utilities.request(http.MethodDelete, "/delete?id=123", nil)

	controller.Delete(rec, req)

	resp := rec.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
