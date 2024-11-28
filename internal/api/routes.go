package api

import (
	"net/http"

	"github.com/lf-silva/fastTrack/internal/api/handlers"
	"github.com/lf-silva/fastTrack/internal/domain"
)

func NewRouter(domain *domain.QuizService) *http.ServeMux {
	router := http.NewServeMux()
	quizHandler := handlers.NewQuizHandler(domain)
	router.Handle("GET /questions", http.HandlerFunc(quizHandler.GetQuestions))
	router.Handle("POST /submit", http.HandlerFunc(quizHandler.SubmitAnswers))

	return router
}
