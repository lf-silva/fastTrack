package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/lf-silva/fastTrack/model"
)

func TestGetQuestions(t *testing.T) {
	store := &StubQuizStore{
		questions: questions,
	}
	server := NewQuizServer(store)

	t.Run("returns questions", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/questions", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var result []model.Question
		err := json.NewDecoder(response.Body).Decode(&result)
		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
		}

		assertStatus(t, response.Code, http.StatusOK)
		if !reflect.DeepEqual(result, questions) {
			t.Errorf("got %v want %v", result, questions)
		}
	})

}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %d want %d", got, want)
	}
}

type StubQuizStore struct {
	questions []model.Question
}

func (s *StubQuizStore) GetQuestions() []model.Question {
	return s.questions
}

var questions = []model.Question{
	{ID: 1, Question: "What is FastTrack model based on?", Answers: []string{"Complexity", "Singularity", "Dangerous", "Fun"}, CorrectAnswer: 1},
	{ID: 2, Question: "When was FastTrack founded?", Answers: []string{"2010", "2013", "2016", "2019"}, CorrectAnswer: 2},
	{ID: 3, Question: "What is FastTrack core product?", Answers: []string{"iGaming CRM ", "Real state CRM", "Financial Services", "E-commerce"}, CorrectAnswer: 0},
	{ID: 4, Question: "What was last FastTrack component launch?", Answers: []string{"Singularity ", "Greco", "Vimeo", "Rewards"}, CorrectAnswer: 3},
}
