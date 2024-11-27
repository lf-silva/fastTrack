package domain

import (
	"errors"

	"github.com/lf-silva/fastTrack/internal/model"
)

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

func (d *QuizDomain) SubmitAnswers(answers []model.Answer) (model.Result, error) {
	if err := validateAnswers(answers); err != nil {
		return model.Result{}, err
	}

	var correctAnswers int
	for _, q := range answers {
		question, ok := d.store.GetQuestion(q.QuestionID)
		if ok && question.IsAnswerCorrect(q.UserAnswer) {
			correctAnswers++
		}
	}
	d.store.SubmitScore(correctAnswers)

	return model.Result{CorrectAnswers: correctAnswers, Score: 0.0}, nil
}

func validateAnswers(answers []model.Answer) error {
	a := make(map[int]int)
	for _, item := range answers {
		_, ok := a[item.QuestionID]
		if ok {
			return errors.New(MoreThanOneAnswerProvided)
		} else {
			a[item.QuestionID] = item.UserAnswer
		}
	}
	return nil
}
