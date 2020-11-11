package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/treemux"
	"net/http"
)

type Controller struct {
	Redis *redis.Client
}

func InitController(red *redis.Client) Controller {
	return Controller{Redis: red}
}

// SHOW DATA
func (_r *Controller) Show(writer http.ResponseWriter, request treemux.Request) error {
	var response Response
	var user User

	ID := fmt.Sprintf("user:%s", request.URL.Query().Get("id"))

	data, errGetData := _r.Redis.HMGet(context.Background(), ID, "ID", "name", "age").Result()

	if errGetData != nil {
		response.Message = errGetData.Error()
		return treemux.JSON(writer, response)
	}

	if data[0] == nil {
		response.Message = "Data not found"
		return treemux.JSON(writer, response)
	}

	user.ID = StrToInt(data[0].(string))
	user.Name = data[1].(string)
	user.Age = StrToInt(data[2].(string))

	response.Data = user
	response.Message = "Success Show Data"

	return treemux.JSON(writer, response)
}

// STORE DATA
func (_r *Controller) Store(writer http.ResponseWriter, request treemux.Request) error {
	var user User
	var response Response

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		response.Message = err.Error()
		return treemux.JSON(writer, response)
	}

	ID := fmt.Sprintf("user:%s", IntToStr(user.ID))

	_r.Redis.HSet(context.Background(), ID, "ID", user.ID, "name", user.Name, "age", user.Age)

	response.Data = nil
	response.Message = "Success Storing Data"

	return treemux.JSON(writer, response)
}

// UPDATE DATA
func (_r *Controller) Update(writer http.ResponseWriter, request treemux.Request) error {
	var user User
	var response Response

	err := json.NewDecoder(request.Body).Decode(&user)
	ID := request.URL.Query().Get("id")

	if err != nil {
		response.Message = err.Error()
		treemux.JSON(writer, response)
	}

	key := fmt.Sprintf("user:%s", ID)

	_r.Redis.HSet(context.Background(), key, "ID", ID, "name", user.Name, "age", user.Age)

	response.Message = "Success update data"
	response.Data = nil

	return treemux.JSON(writer, response)
}

// DELETE DATA
func (_r *Controller) Delete(writer http.ResponseWriter, request treemux.Request) error {
	var response Response

	ID := request.URL.Query().Get("id")

	key := fmt.Sprintf("user:%s", ID)

	_r.Redis.HDel(context.Background(), key, "ID", "name", "age")

	response.Message = "Success Delete Data"
	response.Data = nil

	return treemux.JSON(writer, response)
}
