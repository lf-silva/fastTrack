package main

import (
	"log"
	"net/http"

	"github.com/lf-silva/fastTrack/internal/repo"
	"github.com/lf-silva/fastTrack/internal/server"
)

func main() {
	handler := server.NewQuizHandler(repo.NewInMemoryRepo())
	server := server.NewQuizServer(handler)
	log.Fatal(http.ListenAndServe(":5000", server))
}
