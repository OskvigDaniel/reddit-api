package handlers

import (
	"encoding/json"
	"net/http"
	"reddit-api/internal/store"
)

func HandleUsersByPostCount(us *store.UserStore) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			users := us.TopUsers
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(users)
		},
	)
}
