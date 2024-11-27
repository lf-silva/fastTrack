package main

import (
	"log"
	"net/http"

	"github.com/lf-silva/fastTrack/internal/api"
	"github.com/lf-silva/fastTrack/internal/domain"
	"github.com/lf-silva/fastTrack/internal/repo"
)

func main() {
	domain := domain.NewQuizDomain(repo.NewInMemoryRepo())
	server := api.NewRouter(domain)
	log.Fatal(http.ListenAndServe(":5000", server))
}
