package usercontroller

import (
	"auth-service/internal/server"
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
			response := server.Response{Status: http.StatusInternalServerError, Message: "error", Data: server.Data{"error": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := server.Response{Status: http.StatusOK, Message: "success", Data: server.Data{"users": users}}
		json.NewEncoder(rw).Encode(response)
	}
}
