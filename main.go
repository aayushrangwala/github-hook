package main

import (
	"log"
	"net/http"
	"strconv"

	"githubhook/api"
	"githubhook/util"
)

func main() {
	r := api.NewRouter()
	log.Printf("Listening at %d port...", util.DefaultServicePort)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(util.DefaultServicePort), r))
}
