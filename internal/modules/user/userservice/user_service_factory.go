package userservice

import (
	"github.com/ruffHub/auth-service/internal/modules/user/userrepository"
)

func NewUserService(r userrepository.Repository) Service {
	return Service{userRepository: r}
}
