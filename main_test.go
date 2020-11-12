package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/beinan/fastid"
	"github.com/stretchr/testify/assert"
	"golang-redis/controllers"
	"golang-redis/models"
	"golang-redis/utilities"
	"net/http"
	"testing"
)

var IDs int64

// TEST STORE
func TestStore(t *testing.T) {
	var user models.UserModel
	controller := controllers.InitController(Conn())

	IDs = fastid.CommonConfig.GenInt64ID()

	user.ID = IDs
	user.Username = "Tanthowi"
	user.Email = "bill.tj@icloud.com"
	user.Password = "tanthowi"
	user.RoleID = 1

	out, _ := json.Marshal(user)
	payload := bytes.NewBuffer(out)

	req, rec := utilities.Request(http.MethodPost, "/api/v1/store", payload)

	controller.Store(rec, req)
	resp := rec.Result()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}

// TEST SHOW
func TestShow(t *testing.T) {
	controller := controllers.InitController(Conn())
	req, rec := utilities.Request(http.MethodGet, fmt.Sprintf("/api/v1/show/%s", utilities.IntToStr(int(IDs))), nil)

	controller.Show(rec, req)
	resp := rec.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// TEST UPDATE
func TestUpdate(t *testing.T) {

	var user models.UserModel
	controller := controllers.InitController(Conn())

	user.Email = "bill.tanthowi.j@gmail.com"
	user.Username = "billtanthowi"
	user.Password = "tanthowi"
	user.RoleID = 1

	out, _ := json.Marshal(user)
	payload := bytes.NewBuffer(out)

	req, rec := utilities.Request(http.MethodPost, fmt.Sprintf("/update/%s", utilities.IntToStr(int(IDs))), payload)

	controller.Update(rec, req)
	resp := rec.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// TEST DELETE
func TestDelete(t *testing.T) {
	controller := controllers.InitController(Conn())

	req, rec := utilities.Request(http.MethodDelete, fmt.Sprintf("/delete/%s", utilities.IntToStr(int(IDs))), nil)

	controller.Delete(rec, req)

	resp := rec.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
