package usercontroller

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/ruffHub/auth-service/internal/modules/user/usermodel"
	"github.com/ruffHub/auth-service/internal/server"
	"net/http"
	"time"
)

// CreateUser method implements UserCreator
func (c Controller) CreateUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var payload usermodel.User
		var validate = validator.New()

		//validate the request body
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := server.Response{Status: http.StatusBadRequest, Message: "error", Data: server.Data{"error": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&payload); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := server.Response{Status: http.StatusBadRequest, Message: "error", Data: server.Data{"error": validationErr.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		createdUser, err := c.userService.CreateUser(ctx, payload)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := server.Response{Status: http.StatusInternalServerError, Message: "error", Data: server.Data{"error": err}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		response := server.Response{Status: http.StatusCreated, Message: "success", Data: server.Data{"user": createdUser}}
		json.NewEncoder(rw).Encode(response)
	}
}
