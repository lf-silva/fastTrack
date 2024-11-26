package main

import (
	"log"
	"net/http"

	"github.com/lf-silva/fastTrack/internal/handler"
	"github.com/lf-silva/fastTrack/internal/routes"
	"github.com/lf-silva/fastTrack/internal/store"
)

func main() {
	handler := handler.NewQuizHandler(store.NewInMemoryQuizStore())
	server := routes.NewQuizServer(handler)
	log.Fatal(http.ListenAndServe(":5000", server))
}
