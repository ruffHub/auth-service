package usercontroller

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ruffHub/auth-service/internal/server"
	"net/http"
	"time"
)

// GetUser method implements UserGetter
func (c Controller) GetUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		params := mux.Vars(r)
		userId := params["userId"]

		user, err := c.userService.GetUser(ctx, userId)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := server.Response{Status: http.StatusInternalServerError, Message: "error", Data: server.Data{"error": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusOK)
		response := server.Response{Status: http.StatusOK, Message: "success", Data: server.Data{"user": user}}
		json.NewEncoder(rw).Encode(response)
	}
}
