package svc

import (
	"log"
	"net/http"
	"strings"

	"githubhook/pkg/github"
	"githubhook/util"
)

// Health endpoint returns the healthy string if the svc is reachable
func Health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if _, err := w.Write([]byte("Hi there, I am healthy!!")); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Error encoding or returning the response", err)
	}
	w.WriteHeader(http.StatusOK)
}

// Clone endpoint clones the repository by the name from the request header and using the auth token
func Clone(w http.ResponseWriter, r *http.Request) {
	var clonePath string
	var err error

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	auth := r.Header.Get(util.AuthTokenKey)
	if auth == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal("Invalid request. No Auth token found")
	}

	repo := r.Header.Get(util.RepositoryKey)
	if repo == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal("Invalid request. No repository name found")
	}

	// oAuth2.0 Token will be in the Header of Type [Bearer]
	auth = strings.Split(auth, util.Space)[1]

	c, clErr := github.NewGitClient(auth)
	if clErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("Error creating the client or github repo: %s", clErr.Error())
	}

	if clonePath, err = c.Clone(repo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalf(err.Error())
	}

	if _, err = w.Write([]byte(clonePath)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("Error encoding or returning the response: %s", err.Error())
	}
	w.WriteHeader(http.StatusOK)
}

func Fetch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

func Help(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

func ListRepos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

func Checkout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

func Merge(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}
