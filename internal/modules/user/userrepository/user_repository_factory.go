package userrepository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// NewUserRepository returns initialized Repository
func NewUserRepository(c *mongo.Collection) Repository {
	return Repository{c}
}
