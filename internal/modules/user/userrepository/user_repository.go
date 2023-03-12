package userrepository

import (
	"context"
	"github.com/ruffHub/auth-service/internal/modules/user/usermodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Repository struct {
	collection *mongo.Collection
}

// Create method implements userservice.UserCreator
func (r Repository) Create(ctx context.Context, u usermodel.User) (usermodel.User, error) {
	result, err := r.collection.InsertOne(ctx, u)
	if err != nil {
		return usermodel.User{}, nil
	}

	createdId := result.InsertedID.(primitive.ObjectID).Hex()

	return usermodel.User{Id: createdId}, nil
}

// Get method implements userservice.UserGetter
func (r Repository) Get(ctx context.Context, userId string) (usermodel.User, error) {
	var user usermodel.User

	objId, objectIdErr := primitive.ObjectIDFromHex(userId)
	if objectIdErr != nil {
		return user, objectIdErr
	}

	filter := bson.M{"_id": objId}

	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, err
		}

		panic(err)
	}

	return user, nil
}

// GetAll method implements userservice.UserAllGetter
func (r Repository) GetAll(ctx context.Context) ([]usermodel.User, error) {
	var users []usermodel.User

	filter := bson.M{}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return users, err
		}

		panic(err)
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user usermodel.User
		if err = cursor.Decode(&user); err != nil {
			log.Println("Error while decoding user item")
		} else {
			users = append(users, user)
		}
	}

	return users, err
}
