package usercontroller

import (
	"auth-service/internal/modules/user/userservice"
	"net/http"
)

type Controller struct {
	userService userservice.UserService
}

type ControllerUseCases interface {
	UserCreator
	UserGetter
	UserAllGetter
}

type UserCreator interface {
	CreateUser() http.HandlerFunc
}

type UserGetter interface {
	GetUser() http.HandlerFunc
}

type UserAllGetter interface {
	GetAllUsers() http.HandlerFunc
}
