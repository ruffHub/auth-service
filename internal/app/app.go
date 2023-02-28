package app

import (
	"auth-service/internal/db"
	"auth-service/internal/server"
)

func ApplicationRun() {
	db.ConnectDB()
	server.StartHttpServer()
}
