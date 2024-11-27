package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/lf-silva/fastTrack/internal/api/handlers"
	"github.com/lf-silva/fastTrack/internal/model"
)

func TestQuestions(t *testing.T) {
	store := &StubQuizHandler{}
	handler := handlers.NewQuizHandler(store)

	t.Run("returns questions", func(t *testing.T) {
		want := []model.Question{
			{ID: 1, Question: "What is FastTrack model based on?", Answers: []string{"Complexity", "Singularity", "Dangerous", "Fun"}},
			{ID: 2, Question: "When was FastTrack founded?", Answers: []string{"2010", "2013", "2016", "2019"}},
			{ID: 3, Question: "What is FastTrack core product?", Answers: []string{"iGaming CRM ", "Real state CRM", "Financial Services", "E-commerce"}},
			{ID: 4, Question: "What was last FastTrack component launch?", Answers: []string{"Singularity ", "Greco", "Vimeo", "Rewards"}},
		}
		request, _ := http.NewRequest(http.MethodGet, "/questions", nil)
		response := httptest.NewRecorder()

		handler.GetQuestions(response, request)

		var got []model.Question
		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
		}

		assertStatus(t, response.Code, http.StatusOK)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSubmitAnswers(t *testing.T) {
	store := &StubQuizHandler{}
	handler := handlers.NewQuizHandler(store)

	t.Run("returns bad request with invalid json", func(t *testing.T) {
		invalidJSON := `{"invalid":}`
		request, _ := http.NewRequest(http.MethodPost, "/submit", bytes.NewReader([]byte(invalidJSON)))
		response := httptest.NewRecorder()

		handler.SubmitAnswers(response, request)

		assertValidateAnswersCalls(t, store.validateCalls, 0)
		assertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("returns bad request with empty json", func(t *testing.T) {
		emptyJSON := ``

		request, _ := http.NewRequest(http.MethodPost, "/submit", bytes.NewReader([]byte(emptyJSON)))
		response := httptest.NewRecorder()

		handler.SubmitAnswers(response, request)

		assertValidateAnswersCalls(t, store.validateCalls, 0)
		assertStatus(t, response.Code, http.StatusBadRequest)
	})

	t.Run("calls handler with valid answers", func(t *testing.T) {
		userAnswers := []model.Answer{
			{QuestionID: 1, UserAnswer: 0},
			{QuestionID: 2, UserAnswer: 1},
			{QuestionID: 3, UserAnswer: 2},
			{QuestionID: 4, UserAnswer: 3},
		}
		body, _ := json.Marshal(userAnswers)

		request, _ := http.NewRequest(http.MethodPost, "/submit", bytes.NewReader([]byte(body)))
		response := httptest.NewRecorder()

		handler.SubmitAnswers(response, request)

		assertValidateAnswersCalls(t, store.validateCalls, 1)
		assertStatus(t, response.Code, http.StatusOK)
	})
}

func assertValidateAnswersCalls(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("validate calls are wrong, got %d want %d", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("response status is wrong, got %d want %d", got, want)
	}
}

type StubQuizHandler struct {
	validateCalls int
}

func (s *StubQuizHandler) GetQuestions() []model.Question {
	return []model.Question{
		{ID: 1, Question: "What is FastTrack model based on?", Answers: []string{"Complexity", "Singularity", "Dangerous", "Fun"}, CorrectAnswer: 1},
		{ID: 2, Question: "When was FastTrack founded?", Answers: []string{"2010", "2013", "2016", "2019"}, CorrectAnswer: 2},
		{ID: 3, Question: "What is FastTrack core product?", Answers: []string{"iGaming CRM ", "Real state CRM", "Financial Services", "E-commerce"}, CorrectAnswer: 0},
		{ID: 4, Question: "What was last FastTrack component launch?", Answers: []string{"Singularity ", "Greco", "Vimeo", "Rewards"}, CorrectAnswer: 3},
	}
}

func (s *StubQuizHandler) ValidateAnswers(answers []model.Answer) int {
	s.validateCalls++
	return 0
}
