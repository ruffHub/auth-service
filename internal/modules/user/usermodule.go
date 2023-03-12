package user

import (
	"github.com/gorilla/mux"
	"github.com/ruffHub/auth-service/internal/config"
	"github.com/ruffHub/auth-service/internal/db"
	"github.com/ruffHub/auth-service/internal/modules/user/usercontroller"
	"github.com/ruffHub/auth-service/internal/modules/user/userrepository"
	"github.com/ruffHub/auth-service/internal/modules/user/userservice"
	"go.mongodb.org/mongo-driver/mongo"
)

type Module struct {
	controller usercontroller.ControllerUseCases
}

func NewUserModule(mongoClient *mongo.Client) Module {
	userCollection := db.GetCollection(mongoClient, config.GetEnvVar("MONGO_DB_USERS_COLLECTION_NAME"))
	userRepository := userrepository.NewUserRepository(userCollection)
	userService := userservice.NewUserService(userRepository)
	userController := usercontroller.NewUserController(userService)

	return Module{userController}
}

func (m Module) RegisterRoutes(router *mux.Router) {
	apiV := "/api/" + config.GetEnvVar("API_VERSION")

	router.HandleFunc(apiV+"/user", m.controller.CreateUser()).Methods("POST")
	router.HandleFunc(apiV+"/user/getAll", m.controller.GetAllUsers()).Methods("GET")
	router.HandleFunc(apiV+"/user/{userId}", m.controller.GetUser()).Methods("GET")
	//router.HandleFunc(apiV+"/user/{userId}", m.controller.EditAUser()).Methods("PUT")
	//router.HandleFunc(apiV+"/user/{userId}", m.controller.DeleteAUser()).Methods("DELETE")
}
