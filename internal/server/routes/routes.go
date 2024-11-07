package routes

import (
	"net/http"
	"reddit-api/internal/server/handlers"
	"reddit-api/internal/store"
)

func AddRoutes(
	mux *http.ServeMux,
	userStore *store.UserStore,
	postStore *store.PostStore,
) {
	mux.Handle("/users", handlers.HandleUsersByPostCount(userStore))
	mux.Handle("/posts", handlers.HandlePostsByUpvoteCount(postStore))
	mux.Handle("/", http.NotFoundHandler())
	mux.Handle("/healthz", handlers.HandleHealthz())

}
