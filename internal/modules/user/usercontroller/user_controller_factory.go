package usercontroller

import (
	"auth-service/internal/modules/user/userservice"
)

// NewUserController returns initialized Controller
func NewUserController(c userservice.Service) Controller {
	return Controller{c}
}