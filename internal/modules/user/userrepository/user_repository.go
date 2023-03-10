package userrepository

import (
	"auth-service/internal/modules/user/usermodel"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

type UserRepository interface {
	UserCreator
	UserGetter
	UserAllGetter
}

type UserCreator interface {
	Create(ctx context.Context, u usermodel.User) (*mongo.InsertOneResult, error)
}

type UserGetter interface {
	Get(ctx context.Context, objectId primitive.ObjectID) *mongo.SingleResult
}

type UserAllGetter interface {
	GetAll(ctx context.Context) (cur *mongo.Cursor, err error)
}

// Create method implements UserCreator
func (r Repository) Create(ctx context.Context, u usermodel.User) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(ctx, u)
}

// Get method implements UserGetter
func (r Repository) Get(ctx context.Context, objId primitive.ObjectID) *mongo.SingleResult {
	filter := bson.M{"_id": objId}

	return r.collection.FindOne(ctx, filter)
}

// GetAll method implements UserAllGetter
func (r Repository) GetAll(ctx context.Context) (cur *mongo.Cursor, err error) {
	filter := bson.M{}
	return r.collection.Find(ctx, filter)
}
