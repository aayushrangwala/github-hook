package api

import (
	"githubhook/util"

	"github.com/gorilla/mux"
)

// NewRouter is the function which creates the mux router according to the paths and their handlers
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		handler := util.Logger(route.HandlerFunc, route.Name)

		router.PathPrefix("/github/repos")

		router.
			Path(route.Pattern).
			Methods(route.Method).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
