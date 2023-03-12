package usercontroller

import (
	"github.com/ruffHub/auth-service/internal/modules/user/userservice"
)

// NewUserController returns initialized Controller
func NewUserController(s userservice.Service) Controller {
	return Controller{userService: s}
}
