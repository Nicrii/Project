package controllers

import (
	"encoding/json"
	_ "github.com/Nicrii/Project/users-api/datasources/users_db"
	"github.com/Nicrii/Project/users-api/domain/users"
	"github.com/Nicrii/Project/users-api/services"
	"github.com/Nicrii/Project/users-api/utils/crypto_utils"
	"github.com/Nicrii/Project/users-api/utils/errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId, err := strconv.ParseInt(params["user_id"], 10, 64)

	if err != nil {
		restErr := errors.NewInternalServerError(err.Error())
		jsonValue, _ := json.Marshal(restErr)
		resp.WriteHeader(restErr.Status)
		resp.Write(jsonValue)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.Status)
		resp.Write(jsonValue)
		return
	}

	jsonValue, err := user.Marshall(req.Header.Get("X-Public") == "true")
	if err != nil {
		restErr := errors.NewInternalServerError(err.Error())
		jsonValue, _ := json.Marshal(restErr)
		resp.WriteHeader(restErr.Status)
		resp.Write(jsonValue)
		return
	}
	resp.Write(jsonValue)
}

func CreateUser(resp http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		resp.WriteHeader(restErr.Status)
		jsonErr, _ := json.Marshal(restErr)
		resp.Write(jsonErr)
		return
	}
	var user users.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		resp.WriteHeader(restErr.Status)
		jsonErr, _ := json.Marshal(restErr)
		resp.Write(jsonErr)
		return
	}

	user.Password = crypto_utils.GetMd5(user.Password)
	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		resp.WriteHeader(saveErr.Status)
		jsonErr, _ := json.Marshal(saveErr)
		resp.Write(jsonErr)
		return
	}
	jsonValue, err := result.Marshall(req.Header.Get("X-Public") == "true")
	if err != nil {
		restErr := errors.NewInternalServerError(err.Error())
		jsonValue, _ := json.Marshal(restErr)
		resp.WriteHeader(restErr.Status)
		resp.Write(jsonValue)
		return
	}
	resp.Write(jsonValue)
}

func UpdateUser(resp http.ResponseWriter, req *http.Request) {
	var user users.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		resp.WriteHeader(restErr.Status)
		jsonErr, _ := json.Marshal(restErr)
		resp.Write(jsonErr)
		return
	}

	result, updateErr := services.UserService.UpdateUser(user)
	if updateErr != nil {
		resp.WriteHeader(updateErr.Status)
		jsonErr, _ := json.Marshal(updateErr)
		resp.Write(jsonErr)
		return
	}

	jsonValue, err := result.Marshall(req.Header.Get("X-Public") == "true")
	if err != nil {
		restErr := errors.NewInternalServerError(err.Error())
		jsonValue, _ := json.Marshal(restErr)
		resp.WriteHeader(restErr.Status)
		resp.Write(jsonValue)
		return
	}
	resp.Write(jsonValue)
}

func DeleteUser(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userId, err := strconv.ParseInt(params["user_id"], 10, 64)

	if err != nil {
		restErr := errors.NewInternalServerError(err.Error())
		jsonValue, _ := json.Marshal(restErr)
		resp.WriteHeader(restErr.Status)
		resp.Write(jsonValue)
		return
	}

	apiErr := services.UserService.DeleteUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.Status)
		resp.Write(jsonValue)
		return
	}

	jsoneValue, _ := json.Marshal(map[string]string{"status": "deleted"})
	resp.WriteHeader(http.StatusOK)
	resp.Write(jsoneValue)
}
