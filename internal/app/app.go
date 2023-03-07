package app

import (
	"auth-service/internal/config"
	"auth-service/internal/db"
	"auth-service/internal/modules/user"
	"auth-service/internal/server"
)

func Run() {
	config.LoadEnv()
	mongoClient := db.NewDBConnection()
	router := server.NewRouter()

	userModule := user.NewUserModule(mongoClient)

	routes := []server.RouteHandlersRegisterer{userModule}

	server.RegisterDefaultRoute(router)
	server.RegisterRoutes(router, routes...)
	server.StartHttpServer(router)
}
