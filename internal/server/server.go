package server

import (
	"github.com/gorilla/mux"
	"github.com/ruffHub/auth-service/internal/config"
	"log"
	"net/http"
)

type RouteHandlersRegisterer interface {
	RegisterRoutes(router *mux.Router)
}

func RegisterRoutes(router *mux.Router, modules ...RouteHandlersRegisterer) {
	for _, module := range modules {
		module.RegisterRoutes(router)
	}
}

func NewRouter() *mux.Router {
	return mux.NewRouter()
}

func StartHttpServer(r *mux.Router) {
	listenAndServe(r)
}

func listenAndServe(r *mux.Router) {
	log.Fatal(http.ListenAndServe(":"+config.GetEnvVar("PORT"), r))
}
