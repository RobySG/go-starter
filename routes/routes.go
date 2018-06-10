package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"rest-api/controllers"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func CreateAuthRoutes(mux *mux.Router, ac *controllers.AuthController) {
	mux.HandleFunc("/", homePage)
	auth := mux.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", ac.Register).Methods("POST")
	auth.HandleFunc("/login", ac.Login).Methods("POST")
}

func CreateAPIRoutes(mux *mux.Router, uc *controllers.UserController, jc *controllers.JobController) {
	api := mux.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/users", uc.Lists).Methods("GET")

	mux.HandleFunc("/job", jc.Create)
	mux.HandleFunc("/job/", jc.Job)
	mux.HandleFunc("/job/feed", jc.Feed)
}
