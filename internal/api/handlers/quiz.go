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
		ValidateAnswers(userAnswers []model.Answer) int
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

	correctAnswers := q.domain.ValidateAnswers(answers)
	json.NewEncoder(w).Encode(correctAnswers)
}
