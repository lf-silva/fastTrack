package handler_test

import (
	"reflect"
	"testing"

	"github.com/lf-silva/fastTrack/internal/handler"
	"github.com/lf-silva/fastTrack/internal/model"
)

func TestGetQuestions(t *testing.T) {
	store := &StubQuizStore{
		questions: questions,
	}
	handler := handler.NewQuizHandler(store)

	t.Run("returns questions", func(t *testing.T) {
		got := handler.GetQuestions()

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
		handler := handler.NewQuizHandler(store)

		answers := []model.Answer{
			{QuestionID: 1, UserAnswer: 0},
			{QuestionID: 2, UserAnswer: 0},
			{QuestionID: 3, UserAnswer: 0},
			{QuestionID: 4, UserAnswer: 0},
		}
		want := 1
		got := handler.ValidateAnswers(answers)

		assertAnswers(t, got, want)
		assertSubmitScoreCalls(t, store.submitCalls, 1)
	})
	t.Run("validates questions and returns correct value", func(t *testing.T) {
		store := &StubQuizStore{
			questions: questions,
		}
		handler := handler.NewQuizHandler(store)

		answers := []model.Answer{
			{QuestionID: 1, UserAnswer: 1},
			{QuestionID: 2, UserAnswer: 2},
			{QuestionID: 3, UserAnswer: 0},
			{QuestionID: 4, UserAnswer: 3},
		}
		want := 4
		got := handler.ValidateAnswers(answers)

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
	questions   map[int]model.Question
	submitCalls int
}

func (s *StubQuizStore) GetQuestions() []model.Question {
	array := make([]model.Question, 0, len(questions))
	for _, q := range questions {
		array = append(array, q)
	}
	return array
}

func (s *StubQuizStore) GetQuestion(id int) (model.Question, bool) {
	q, ok := questions[id]
	return q, ok
}

func (s *StubQuizStore) SubmitScore(score int) {
	s.submitCalls++
}

var questions = map[int]model.Question{
	1: {ID: 1, Question: "What is FastTrack model based on?", Answers: []string{"Complexity", "Singularity", "Dangerous", "Fun"}, CorrectAnswer: 1},
	2: {ID: 2, Question: "When was FastTrack founded?", Answers: []string{"2010", "2013", "2016", "2019"}, CorrectAnswer: 2},
	3: {ID: 3, Question: "What is FastTrack core product?", Answers: []string{"iGaming CRM ", "Real state CRM", "Financial Services", "E-commerce"}, CorrectAnswer: 0},
	4: {ID: 4, Question: "What was last FastTrack component launch?", Answers: []string{"Singularity ", "Greco", "Vimeo", "Rewards"}, CorrectAnswer: 3},
}
