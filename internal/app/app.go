package app

import (
	"github.com/ruffHub/auth-service/internal/config"
	"github.com/ruffHub/auth-service/internal/db"
	"github.com/ruffHub/auth-service/internal/modules/user"
	"github.com/ruffHub/auth-service/internal/server"
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
