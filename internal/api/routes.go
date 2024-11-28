package api

import (
	"net/http"

	"github.com/lf-silva/fastTrack/internal/api/handlers"
	"github.com/lf-silva/fastTrack/internal/api/middleware"
	"github.com/lf-silva/fastTrack/internal/domain"
)

func NewRouter(domain *domain.QuizService) *http.ServeMux {
	router := http.NewServeMux()
	quizHandler := handlers.NewQuizHandler(domain)
	router.Handle("GET /questions", middleware.Logging(http.HandlerFunc(quizHandler.GetQuestions)))
	router.Handle("POST /submit", middleware.Logging(http.HandlerFunc(quizHandler.SubmitAnswers)))

	return router
}
