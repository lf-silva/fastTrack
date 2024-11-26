package server

import "github.com/lf-silva/fastTrack/internal/model"

type QuizStore interface {
	GetQuestions() []model.Question
	GetQuestion(id int) (model.Question, bool)
	SubmitScore(result int)
}
type QuizHandler struct {
	store QuizStore
}

func NewQuizHandler(store QuizStore) *QuizHandler {
	return &QuizHandler{
		store: store,
	}
}
func (h *QuizHandler) GetQuestions() []model.Question {
	return h.store.GetQuestions()
}

func (h *QuizHandler) ValidateAnswers(answers []model.Answer) int {
	var correctAnswers int
	for _, q := range answers {
		question, ok := h.store.GetQuestion(q.QuestionID)
		if ok && question.IsAnswerCorrect(q.UserAnswer) {
			correctAnswers++
		}
	}
	h.store.SubmitScore(correctAnswers)
	return correctAnswers
}
