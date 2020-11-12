package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/beinan/fastid"
	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/treemux"
	"golang-redis/models"
	"golang-redis/presenter"
	"golang-redis/utilities"
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
	var response presenter.Response
	var user models.UserModel

	ID := fmt.Sprintf("user:%s", request.Param("id"))

	data, errGetData := _r.Redis.HMGet(context.Background(), ID, "id", "email", "username", "password", "role_id").Result()

	if errGetData != nil {
		response.Message = errGetData.Error()
		return treemux.JSON(writer, response)
	}

	if data[0] == nil {
		response.Message = "Data not found"
		return treemux.JSON(writer, response)
	}

	user.ID = int64(utilities.StrToInt(data[0].(string)))
	user.Email = data[1].(string)
	user.Username = data[2].(string)
	user.Password = data[3].(string)
	user.RoleID = utilities.StrToInt(data[3].(string))

	response.Data = user
	response.Message = "Success Show Data"

	return treemux.JSON(writer, response)
}

// STORE DATA
func (_r *Controller) Store(writer http.ResponseWriter, request treemux.Request) error {
	var user models.UserModel
	var response presenter.Response

	IDs := fastid.CommonConfig.GenInt64ID()

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		response.Message = err.Error()
		return treemux.JSON(writer, response)
	}

	ID := fmt.Sprintf("user:%s", utilities.IntToStr(int(user.ID)))

	if user.ID == 0 {
		ID = fmt.Sprintf("user:%s", utilities.IntToStr(int(IDs)))
		user.ID = IDs
	}

	_r.Redis.HSet(context.Background(), ID, "id", IDs, "email", user.Email, "username", user.Username, "password", user.Password, "role_id", user.RoleID)

	response.Data = user
	response.Message = "Success Storing Data"

	writer.WriteHeader(http.StatusCreated)
	return treemux.JSON(writer, response)
}

// UPDATE DATA
func (_r *Controller) Update(writer http.ResponseWriter, request treemux.Request) error {
	var user models.UserModel
	var response presenter.Response

	err := json.NewDecoder(request.Body).Decode(&user)
	ID := request.Param("id")

	if err != nil {
		response.Message = err.Error()
		treemux.JSON(writer, response)
	}

	key := fmt.Sprintf("user:%s", ID)

	_r.Redis.HSet(context.Background(), key, "id", ID, "email", user.Email, "username", user.Username, "password", user.Password, "role_id", user.RoleID)

	response.Message = "Success update data"
	response.Data = nil

	return treemux.JSON(writer, response)
}

// DELETE DATA
func (_r *Controller) Delete(writer http.ResponseWriter, request treemux.Request) error {
	var response presenter.Response

	ID := request.Param("id")

	key := fmt.Sprintf("user:%s", ID)

	_r.Redis.HDel(context.Background(), key, "id", "email", "username", "password", "role_id")

	response.Message = "Success Delete Data"
	response.Data = nil

	return treemux.JSON(writer, response)
}
