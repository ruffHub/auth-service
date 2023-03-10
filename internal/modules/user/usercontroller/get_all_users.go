package usercontroller

import (
	"auth-service/internal/responses"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// GetAllUsers
//method implements UserAllGetter
func (c Controller) GetAllUsers() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		users, err := c.userService.GetAllUsers(ctx)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: responses.Data{"error": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := responses.Response{Status: http.StatusOK, Message: "success", Data: responses.Data{"users": users}}
		json.NewEncoder(rw).Encode(response)
	}
}
