package domain_test

import (
	"reflect"
	"testing"

	"github.com/lf-silva/fastTrack/internal/domain"
	"github.com/lf-silva/fastTrack/internal/model"
)

func TestGetQuestions(t *testing.T) {
	store := &StubQuizStore{
		questions: questions,
	}
	domain := domain.NewQuizDomain(store)

	t.Run("returns questions", func(t *testing.T) {
		got := domain.GetQuestions()

		if !reflect.DeepEqual(got, questions) {
			t.Errorf("got %v want %v", got, questions)
		}
	})
}

func TestValidateAnswers(t *testing.T) {
	t.Run("validates questions and returns correct value", func(t *testing.T) {
		store := &StubQuizStore{
			questions: questions,
		}
		domain := domain.NewQuizDomain(store)

		answers := []model.Answer{
			{QuestionID: 1, UserAnswer: 0},
			{QuestionID: 2, UserAnswer: 0},
			{QuestionID: 3, UserAnswer: 0},
			{QuestionID: 4, UserAnswer: 0},
		}
		want := 1
		got := domain.ValidateAnswers(answers)

		assertAnswers(t, got, want)
		assertSubmitScoreCalls(t, store.submitCalls, 1)
	})
	t.Run("validates questions and returns correct value", func(t *testing.T) {
		store := &StubQuizStore{
			questions: questions,
		}
		domain := domain.NewQuizDomain(store)

		answers := []model.Answer{
			{QuestionID: 1, UserAnswer: 1},
			{QuestionID: 2, UserAnswer: 2},
			{QuestionID: 3, UserAnswer: 0},
			{QuestionID: 4, UserAnswer: 3},
		}
		want := 4
		got := domain.ValidateAnswers(answers)

		assertAnswers(t, got, want)
		assertSubmitScoreCalls(t, store.submitCalls, 1)
	})
}

func assertAnswers(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("correct answers are wrong, got %d want %d", got, want)
	}
}
func assertSubmitScoreCalls(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("calls to submit result are wrong, got %d want %d", got, want)
	}
}

type StubQuizStore struct {
	questions   []model.Question
	submitCalls int
}

func (s *StubQuizStore) GetQuestions() []model.Question {
	return questions
}

func (s *StubQuizStore) GetQuestion(id int) (model.Question, bool) {
	for _, q := range questions {
		if q.ID == id {
			return q, true
		}
	}
	return model.Question{}, false
}

func (s *StubQuizStore) SubmitScore(score int) {
	s.submitCalls++
}

var questions = []model.Question{
	{ID: 1, Question: "What is FastTrack model based on?", Answers: []string{"Complexity", "Singularity", "Dangerous", "Fun"}, CorrectAnswer: 1},
	{ID: 2, Question: "When was FastTrack founded?", Answers: []string{"2010", "2013", "2016", "2019"}, CorrectAnswer: 2},
	{ID: 3, Question: "What is FastTrack core product?", Answers: []string{"iGaming CRM ", "Real state CRM", "Financial Services", "E-commerce"}, CorrectAnswer: 0},
	{ID: 4, Question: "What was last FastTrack component launch?", Answers: []string{"Singularity ", "Greco", "Vimeo", "Rewards"}, CorrectAnswer: 3},
}
