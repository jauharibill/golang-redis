package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

// LISTING DATA
func Index(writer http.ResponseWriter, request *http.Request) {
	var response Response

	name := request.URL.Query().Get("name")

	response.Message = fmt.Sprintf("Success get Data %s", name)
	response.Data = nil

	out, _ := json.Marshal(response)

	writer.Write(out)
	writer.WriteHeader(http.StatusOK)
}

// STORE DATA
func Store(writer http.ResponseWriter, request *http.Request) {
	var user User
	var response Response

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	red := Conn()
	red.Set(context.Background(), "name", user.Name, 0).Err()

	// HMSET
	red.HMSet(context.Background(), "user:123", "name", "bill", "age", 26)
	red.HMSet(context.Background(), "user:125", "name", "tanthowi", "age", 27)
	red.HMSet(context.Background(), "user:124", "name", "jauhari", "age", 28)

	// HMGET
	user123, _ := red.HMGet(context.Background(), "user:123", "name", "age").Result()
	user124, _ := red.HMGet(context.Background(), "user:124", "name", "age").Result()
	user125, _ := red.HMGet(context.Background(), "user:125", "name", "age").Result()

	// OUTPUT REDIS
	log.Info().Msg(fmt.Sprintf("my name is %s, I am %d years old", user123[0].(string), StrToInt(user123[1].(string))))
	log.Info().Msg(fmt.Sprintf("my name is %s, I am %d years old", user124[0].(string), StrToInt(user124[1].(string))))
	log.Info().Msg(fmt.Sprintf("my name is %s, I am %d years old", user125[0].(string), StrToInt(user125[1].(string))))

	response.Data = user
	response.Message = "Success Storing Data"

	res, _ := json.Marshal(response)

	writer.Write(res)
}

// UPDATE DATA
func Update(writer http.ResponseWriter, request *http.Request) {
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
func Delete(writer http.ResponseWriter, request *http.Request) {
	var response Response

	ID := request.URL.Query().Get("id")

	response.Message = fmt.Sprintf("Success Delete Data %s", ID)
	response.Data = nil

	res, _ := json.Marshal(response)

	writer.Write(res)
}
