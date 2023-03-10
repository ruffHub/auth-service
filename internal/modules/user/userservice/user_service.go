package userservice

import (
	"auth-service/internal/modules/user/usermodel"
	"auth-service/internal/modules/user/userrepository"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Service struct {
	userRepository userrepository.UserRepository
}

type UserService interface {
	CreateUser(ctx context.Context, user usermodel.User) (usermodel.User, error)
	GetUser(ctx context.Context, userId string) (usermodel.User, error)
	GetAllUsers(ctx context.Context) ([]usermodel.User, error)
}

// CreateUser method implements UserService.GetUser
func (s Service) CreateUser(ctx context.Context, user usermodel.User) (usermodel.User, error) {
	var createdUser usermodel.User
	result, err := s.userRepository.Create(ctx, user)

	if err != nil {
		return createdUser, nil
	}

	return usermodel.User{Id: result.InsertedID.(primitive.ObjectID).Hex()}, nil
}

// GetUser method implements UserService.GetUser
func (s Service) GetUser(ctx context.Context, userId string) (usermodel.User, error) {
	objId, _ := primitive.ObjectIDFromHex(userId)
	var user usermodel.User

	err := s.userRepository.Get(ctx, objId).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, err
		}

		panic(err)
	}

	return user, err
}

// GetAllUsers method implements UserService.GetAllUsers
func (s Service) GetAllUsers(ctx context.Context) ([]usermodel.User, error) {
	var users []usermodel.User

	cursor, err := s.userRepository.GetAll(ctx)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return users, err
		}

		panic(err)
	}

	//reading from the db in an optimal way
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var singleUser usermodel.User
		if err = cursor.Decode(&singleUser); err != nil {
			log.Println("Error while decoding user item")
		} else {
			users = append(users, singleUser)
		}
	}

	return users, err
}
