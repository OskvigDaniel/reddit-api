package server

import (
	"net/http"
	"reddit-api/internal/server/routes"
	"reddit-api/internal/store"
)

func CreateServer(
	userStore *store.UserStore,
	postStore *store.PostStore,
) http.Handler {
	mux := http.NewServeMux()

	routes.AddRoutes(mux, userStore, postStore)
	var handler http.Handler = mux

	// middleware goes here

	return handler
}
