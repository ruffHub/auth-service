package usercontroller

import (
	"auth-service/internal/modules/user/usermodel"
	"auth-service/internal/responses"
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
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
			response := responses.Response{Status: http.StatusBadRequest, Message: "error", Data: responses.Data{"error": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&payload); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.Response{Status: http.StatusBadRequest, Message: "error", Data: responses.Data{"error": validationErr.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		createdUser, err := c.userService.CreateUser(ctx, payload)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: responses.Data{"error": err}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		response := responses.Response{Status: http.StatusCreated, Message: "success", Data: responses.Data{"user": createdUser}}
		json.NewEncoder(rw).Encode(response)
	}
}
