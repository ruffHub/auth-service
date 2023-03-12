package userservice

import (
	"context"
	"github.com/ruffHub/auth-service/internal/modules/user/usermodel"
)

type UserRepository interface {
	UserCreator
	UserGetter
	UserAllGetter
}

type UserCreator interface {
	Create(ctx context.Context, u usermodel.User) (usermodel.User, error)
}

type UserGetter interface {
	Get(ctx context.Context, userId string) (usermodel.User, error)
}

type UserAllGetter interface {
	GetAll(ctx context.Context) ([]usermodel.User, error)
}

type Service struct {
	userRepository UserRepository
}

// CreateUser method implements UserService.GetUser
func (s Service) CreateUser(ctx context.Context, user usermodel.User) (usermodel.User, error) {
	createdUser, err := s.userRepository.Create(ctx, user)

	return createdUser, err
}

// GetUser method implements UserService.GetUser
func (s Service) GetUser(ctx context.Context, userId string) (usermodel.User, error) {
	user, err := s.userRepository.Get(ctx, userId)

	return user, err
}

// GetAllUsers method implements UserService.GetAllUsers
func (s Service) GetAllUsers(ctx context.Context) ([]usermodel.User, error) {
	users, err := s.userRepository.GetAll(ctx)

	return users, err
}
