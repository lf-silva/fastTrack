package api

import (
	"encoding/json"
	"net/http"

	"github.com/lf-silva/fastTrack/model"
)

const jsonContentType = "application/json"

type QuizStore interface {
	GetQuestions() []model.Question
}

type QuizServer struct {
	store QuizStore
	http.Handler
}

func NewQuizServer(store QuizStore) *QuizServer {
	q := new(QuizServer)

	q.store = store
	router := http.NewServeMux()
	router.Handle("/questions", http.HandlerFunc(q.quizHandler))

	q.Handler = router
	return q
}

func (q *QuizServer) quizHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(q.store.GetQuestions())
}
