package server

import (
	"auth-service/internal/config"
	"auth-service/internal/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartHttpServer() {
	router := newRouter()
	registerRoutes(router)
	listenAndServe(router)
}

func newRouter() *mux.Router {
	router := mux.NewRouter()

	return router
}

func registerRoutes(r *mux.Router) {
	routes.RegisterUserRoutes(r)
}

func listenAndServe(r *mux.Router) {
	log.Fatal(http.ListenAndServe(":"+config.GetEnvVar("PORT"), r))
}
