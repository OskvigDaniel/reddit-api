package handlers

import (
	"encoding/json"
	"net/http"
	"reddit-api/internal/store"
)

func HandlePostsByUpvoteCount(ps *store.PostStore) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			posts := ps.TopPosts
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(posts)
		},
	)
}
