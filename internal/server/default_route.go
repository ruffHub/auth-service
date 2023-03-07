package server

import (
	"auth-service/internal/responses"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterDefaultRoute(router *mux.Router) {
	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(rw).Encode(responses.Data{"data": "Server is up"})
		if err != nil {
			fmt.Println("Err: ", err)
			return
		}
	})
}
