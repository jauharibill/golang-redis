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
}

func (_r *UserHandler) Show(writer http.ResponseWriter, request *http.Request) {
	var response Response
	var user User

	//ID := request.URL.Query().Get("id")

	data, _ := _r.Redis.HMGet(context.Background(), "user:123", "ID", "name", "age").Result()

	user.ID = StrToInt(data[0].(string))
	user.Name = data[1].(string)
	user.Age = StrToInt(data[2].(string))

	response.Data = user
	response.Message = "Success Show Data"

	out, _ := json.Marshal(response)

	writer.Write(out)
	writer.WriteHeader(http.StatusOK)
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

	fmt.Println(fmt.Sprintf("user:%s", IntToStr(user.ID)))

	_r.Redis.HSet(context.Background(), "user:123", "name", "bill", "age", 25)

	response.Data = user
	response.Message = "Success Storing Data"

	res, _ := json.Marshal(response)

	writer.Write(res)
	writer.WriteHeader(http.StatusCreated)
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

	response.Message = fmt.Sprintf("Success update data %s", ID)
	response.Data = user

	res, _ := json.Marshal(response)

	writer.Write(res)
}

// DELETE DATA
func (_r *UserHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	var response Response

	ID := request.URL.Query().Get("id")

	response.Message = fmt.Sprintf("Success Delete Data %s", ID)
	response.Data = nil

	res, _ := json.Marshal(response)

	writer.Write(res)
}
