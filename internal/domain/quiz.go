package domain

import (
	"errors"

	"github.com/lf-silva/fastTrack/internal/model"
)

type QuizStore interface {
	GetQuestions() []model.Question
	GetQuestion(id int) (model.Question, bool)
	SubmitScore(result int)
	GetIndexes(score int) (int, int)
}
type QuizService struct {
	store QuizStore
}

func NewQuizService(store QuizStore) *QuizService {
	return &QuizService{
		store: store,
	}
}
func (d *QuizService) GetQuestions() []model.Question {
	return d.store.GetQuestions()
}

func (d *QuizService) SubmitAnswers(answers []model.Answer) (model.Result, error) {
	if err := ensureHasOnlyOneAnswerPerQuestion(answers); err != nil {
		return model.Result{}, err
	}

	correctAnswers := validateAnswers(answers, d)
	d.store.SubmitScore(correctAnswers)

	relativeScore := calculateRelativeScore(correctAnswers, d)

	return model.Result{CorrectAnswers: correctAnswers, Score: relativeScore}, nil
}

func calculateRelativeScore(correctAnswers int, d *QuizService) float64 {
	userIndex, totalScores := d.store.GetIndexes(correctAnswers)
	return float64(userIndex) / float64(totalScores)
}

func validateAnswers(answers []model.Answer, d *QuizService) int {
	var correctAnswers int
	for _, q := range answers {
		question, ok := d.store.GetQuestion(q.QuestionID)
		if ok && question.IsAnswerCorrect(q.UserAnswer) {
			correctAnswers++
		}
	}
	return correctAnswers
}

func ensureHasOnlyOneAnswerPerQuestion(answers []model.Answer) error {
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
