package main

import (
	"log"
	"net/http"

	"github.com/lf-silva/fastTrack/api"
	"github.com/lf-silva/fastTrack/store"
)

func main() {
	server := api.NewQuizServer(store.NewInMemoryQuizStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}
