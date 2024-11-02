package routes

import (
	"github.com/gorilla/mux"
	"github.com/yusufborucu/go-jwt-auth/controllers"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/profile", controllers.Profile).Methods("GET")
}
