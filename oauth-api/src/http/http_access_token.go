package http

import (
	"encoding/json"
	"github.com/Nicrii/Project/oauth-api/src/domain/access_token"
	"github.com/Nicrii/Project/oauth-api/src/utils/errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type AccessTokenHandler interface {
	GetById(resp http.ResponseWriter, req *http.Request)
	Create(resp http.ResponseWriter, req *http.Request)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(resp http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	atId, err := strconv.ParseInt(params["access_token_id"], 10, 64)

	if err != nil {
		restErr := errors.NewInternalServerError(err.Error())
		jsonValue, _ := json.Marshal(restErr)
		resp.WriteHeader(restErr.Status)
		resp.Write(jsonValue)
		return
	}
	resp.WriteHeader(http.StatusOK)
	atIdJson, _ := json.Marshal(atId)
	resp.Write(atIdJson)
}

func (handler *accessTokenHandler) Create(resp http.ResponseWriter, req *http.Request) {
	var at access_token.AccessToken
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		resp.WriteHeader(restErr.Status)
		jsonErr, _ := json.Marshal(restErr)
		resp.Write(jsonErr)
		return
	}
	err = json.Unmarshal(body, &at)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		resp.WriteHeader(restErr.Status)
		jsonErr, _ := json.Marshal(restErr)
		resp.Write(jsonErr)
		return
	}

	if err := handler.service.Create(at); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		resp.WriteHeader(restErr.Status)
		jsonErr, _ := json.Marshal(restErr)
		resp.Write(jsonErr)
		return
	}

	resp.WriteHeader(http.StatusCreated)
	atJson, _ := json.Marshal(at)
	resp.Write(atJson)
}
