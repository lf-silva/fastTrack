package routes

import (
	"encoding/json"
	"net/http"

	"github.com/lf-silva/fastTrack/internal/model"
)

const jsonContentType = "application/json"

type QuizHandler interface {
	GetQuestions() []model.Question
	ValidateAnswers([]model.Answer) int
}

type QuizServer struct {
	handler QuizHandler
	http.Handler
}

func NewQuizServer(handler QuizHandler) *QuizServer {
	q := new(QuizServer)

	q.handler = handler
	router := http.NewServeMux()
	router.Handle("GET /questions", http.HandlerFunc(q.getQuestions))
	router.Handle("POST /submit", http.HandlerFunc(q.validateAnswers))

	q.Handler = router
	return q
}

func (q *QuizServer) getQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(q.handler.GetQuestions())
}

func (q *QuizServer) validateAnswers(w http.ResponseWriter, r *http.Request) {
	var answers []model.Answer
	if err := json.NewDecoder(r.Body).Decode(&answers); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	correctAnswers := q.handler.ValidateAnswers(answers)
	json.NewEncoder(w).Encode(correctAnswers)
}
