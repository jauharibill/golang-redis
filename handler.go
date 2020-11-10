package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net/http"
)

type UserHandler struct {
	Redis *redis.Client
}

func NewHandler(red *redis.Client) UserHandler {
	return UserHandler{Redis: red}
}

// LISTING DATA
func (_r *UserHandler) Index(writer http.ResponseWriter, request *http.Request) {
	var response Response

	response.Message = fmt.Sprintf("Success get Data")
	response.Data = nil

	out, _ := json.Marshal(response)

	writer.Write(out)
	writer.WriteHeader(http.StatusOK)
	return
}

func (_r *UserHandler) Show(writer http.ResponseWriter, request *http.Request) {
	var response Response
	var user User

	ID := fmt.Sprintf("user:%s", request.URL.Query().Get("id"))

	data, errGetData := _r.Redis.HMGet(context.Background(), ID, "name", "age").Result()

	if errGetData != nil {
		response.Message = errGetData.Error()
		out, _ := json.Marshal(response)
		writer.Write(out)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if data[0] == nil {
		response.Message = "Data not found"
		out, _ := json.Marshal(response)

		writer.Write(out)
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	user.Name = data[0].(string)
	user.Age = StrToInt(data[1].(string))

	response.Data = user
	response.Message = "Success Show Data"

	out, _ := json.Marshal(response)

	writer.Write(out)
	writer.WriteHeader(http.StatusOK)
	return
}

// STORE DATA
func (_r *UserHandler) Store(writer http.ResponseWriter, request *http.Request) {
	var user User
	var response Response

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	ID := fmt.Sprintf("user:%s", IntToStr(user.ID))

	_r.Redis.HSet(context.Background(), ID, "name", user.Name, "age", user.Age)

	response.Data = nil
	response.Message = "Success Storing Data"

	res, _ := json.Marshal(response)

	writer.Write(res)
	writer.WriteHeader(http.StatusCreated)
	return
}

// UPDATE DATA
func (_r *UserHandler) Update(writer http.ResponseWriter, request *http.Request) {
	var user User
	var response Response

	err := json.NewDecoder(request.Body).Decode(&user)
	ID := request.URL.Query().Get("id")

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	key := fmt.Sprintf("user:%s", ID)

	_r.Redis.HSet(context.Background(), key, "name", user.Name, "age", user.Age)

	response.Message = "Success update data"
	response.Data = nil

	res, _ := json.Marshal(response)

	writer.Write(res)
	return
}

// DELETE DATA
func (_r *UserHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	var response Response

	ID := request.URL.Query().Get("id")

	key := fmt.Sprintf("user:%s", ID)

	_r.Redis.HDel(context.Background(), key, "name", "age")

	response.Message = "Success Delete Data"
	response.Data = nil

	res, _ := json.Marshal(response)

	writer.Write(res)
	return
}
