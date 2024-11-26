package server

import (
	"encoding/json"
	"net/http"

	"github.com/lf-silva/fastTrack/internal/model"
)

const jsonContentType = "application/json"

type Domain interface {
	GetQuestions() []model.Question
	ValidateAnswers([]model.Answer) int
}

type QuizServer struct {
	domain Domain
	http.Handler
}

func NewQuizServer(quizDomain Domain) *QuizServer {
	q := new(QuizServer)

	q.domain = quizDomain
	router := http.NewServeMux()
	router.Handle("GET /questions", http.HandlerFunc(q.questions))
	router.Handle("POST /submit", http.HandlerFunc(q.submitAnswers))

	q.Handler = router
	return q
}

func (q *QuizServer) questions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(q.domain.GetQuestions())
}

func (q *QuizServer) submitAnswers(w http.ResponseWriter, r *http.Request) {
	var answers []model.Answer
	if err := json.NewDecoder(r.Body).Decode(&answers); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	correctAnswers := q.domain.ValidateAnswers(answers)
	json.NewEncoder(w).Encode(correctAnswers)
}
