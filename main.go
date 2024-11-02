package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yusufborucu/go-jwt-auth/configs"
	"github.com/yusufborucu/go-jwt-auth/models"
	"github.com/yusufborucu/go-jwt-auth/routes"
)

func main() {
	configs.InitDB()
	models.Migrate(configs.DB)

	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	http.ListenAndServe(":8080", r)
}
