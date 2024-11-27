package domain

import "github.com/lf-silva/fastTrack/internal/model"

type QuizStore interface {
	GetQuestions() []model.Question
	GetQuestion(id int) (model.Question, bool)
	SubmitScore(result int)
}
type QuizDomain struct {
	store QuizStore
}

func NewQuizDomain(store QuizStore) *QuizDomain {
	return &QuizDomain{
		store: store,
	}
}
func (d *QuizDomain) GetQuestions() []model.Question {
	return d.store.GetQuestions()
}

func (d *QuizDomain) ValidateAnswers(answers []model.Answer) int {
	var correctAnswers int
	for _, q := range answers {
		question, ok := d.store.GetQuestion(q.QuestionID)
		if ok && question.IsAnswerCorrect(q.UserAnswer) {
			correctAnswers++
		}
	}
	d.store.SubmitScore(correctAnswers)
	return correctAnswers
}
