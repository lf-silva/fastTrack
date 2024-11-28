package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lf-silva/fastTrack/internal/model"
)

const jsonContentType = "application/json"

type (
	QuizDomain interface {
		GetQuestions() []model.Question
		SubmitAnswers(userAnswers []model.Answer) (model.Result, error)
	}

	QuizHandler struct {
		domain QuizDomain
	}
)

func NewQuizHandler(domain QuizDomain) *QuizHandler {
	return &QuizHandler{domain: domain}
}
func (q *QuizHandler) GetQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(q.domain.GetQuestions())
}

func (q *QuizHandler) SubmitAnswers(w http.ResponseWriter, r *http.Request) {
	var answers []model.Answer
	if err := json.NewDecoder(r.Body).Decode(&answers); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	correctAnswers, err := q.domain.SubmitAnswers(answers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(correctAnswers)
}
