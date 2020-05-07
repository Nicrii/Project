package app

import (
	"fmt"
	"github.com/Nicrii/Project/users-api/controllers"
	"github.com/Nicrii/Project/users-api/logger"
	"github.com/gorilla/mux"
	"net/http"
)

func StartApp() {

	mapUrl()
	router := mux.NewRouter()
	router.HandleFunc("/user/{user_id:[0-9]+}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/create", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/update", controllers.UpdateUser).Methods("POST")
	router.HandleFunc("/user/{user_id:[0-9]+}", controllers.DeleteUser).Methods("DELETE")
	router.Use(loggingPath)

	logger.Info("about to start the application...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}

func loggingPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		fmt.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
func Info() {

}
