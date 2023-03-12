package usercontroller

import (
	"context"
	"github.com/ruffHub/auth-service/internal/modules/user/usermodel"
)

type UserService interface {
	CreateUser(ctx context.Context, user usermodel.User) (usermodel.User, error)
	GetUser(ctx context.Context, userId string) (usermodel.User, error)
	GetAllUsers(ctx context.Context) ([]usermodel.User, error)
}

type Controller struct {
	userService UserService
}
