package userrepository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// NewUserRepository returns initialized Repository
func NewUserRepository(c *mongo.Collection) Repository {
	if c == nil {
		log.Fatal("Error during creation of UserRepository. Collection is nil")
	}

	return Repository{collection: c}
}
