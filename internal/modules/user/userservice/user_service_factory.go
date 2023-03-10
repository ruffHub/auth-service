package userservice

import (
	"auth-service/internal/modules/user/userrepository"
)

func NewUserService(r userrepository.Repository) Service {
	return Service{r}
}
