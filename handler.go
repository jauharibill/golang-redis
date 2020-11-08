package main

import (
	"encoding/json"
	"fmt"
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
