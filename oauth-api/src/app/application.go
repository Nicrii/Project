package app

import (
	"github.com/Nicrii/Project/oauth-api/src/domain/access_token"
	httpService "github.com/Nicrii/Project/oauth-api/src/http"
	"github.com/Nicrii/Project/oauth-api/src/repository/db"
	"github.com/gorilla/mux"
	"net/http"
)

func StartApplication() {
	router := mux.NewRouter()
	atHandler := httpService.NewAccessTokenHandler(access_token.NewService(db.NewRepository()))

	router.HandleFunc("/access_token/{access_token_id:[0-9]+}", atHandler.GetById).Methods("GET")
	router.HandleFunc("/oauth/access_token", atHandler.Create).Methods("POST")

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}

}
