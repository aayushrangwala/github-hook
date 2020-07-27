package api

import (
	"net/http"

	"githubhook/pkg/svc"
)

// Route struct defines the http route mapping
type Route struct {
	// Name of the Route
	Name string

	// Method is the http method to be used for that route
	Method string

	// Pattern is the route path pattern for the endpoint
	Pattern string

	// HandlerFunc is the http handler function for the endpoint route
	HandlerFunc http.HandlerFunc
}

// Routes is the list of all the routes with all the paths and its related handler functions and pattern
var routes = []Route{
	// This Route with the name Health represents the path, method and handler for a route
	{
		Name:        "Health",
		Method:      http.MethodGet,
		Pattern:     "/healthz",
		HandlerFunc: svc.Health,
	},

	// This Route with the name Cloner represents the path, method and handler for a git repo cloning route
	{
		Name:        "Cloner",
		Method:      http.MethodGet,
		Pattern:     "/clone",
		HandlerFunc: svc.Clone,
	},

	// This Route with the name Fetcher represents the path, method and handler for a git repo fetching endpoint
	{
		Name:        "Fetcher",
		Method:      http.MethodGet,
		Pattern:     "/fetch",
		HandlerFunc: svc.Fetch,
	},

	// This Route with the name Checkout represents the path, method and handler for a git repo checkout branch endpoint
	{
		Name:        "Checkout",
		Method:      http.MethodGet,
		Pattern:     "/checkout",
		HandlerFunc: svc.Checkout,
	},

	// This Route with the name Merger represents the path, method and handler for a git branch merge branch
	// from current to target branch endpoint
	{
		Name:        "Merger",
		Method:      http.MethodGet,
		Pattern:     "/merge",
		HandlerFunc: svc.Merge,
	},

	// This Route with the name ListRepo represents the path, method and handler for a git branch merge endpoint
	// Returns the list of repositories for the authenticated user
	{
		Name:        "ListRepo",
		Method:      http.MethodGet,
		Pattern:     "/list",
		HandlerFunc: svc.ListRepos,
	},

	// This Route with the name Helper represents the path, method and handler for the root endpoint
	// which will return with the helping comments
	// Need to keep it at the last to be matched
	{
		Name:        "Helper",
		Method:      http.MethodGet,
		Pattern:     "/help",
		HandlerFunc: svc.Help,
	},

	// This Route with the name Helper represents the path, method and handler for the root endpoint
	// which will return with the helping comments
	// Need to keep it at the last to be matched
	{
		Name:        "Helper",
		Method:      http.MethodGet,
		Pattern:     "/",
		HandlerFunc: svc.Help,
	},
}
