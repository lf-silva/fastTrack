package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/lf-silva/fastTrack/internal/domain"
	"github.com/lf-silva/fastTrack/internal/model"
	"github.com/lf-silva/fastTrack/internal/repo"
)

func TestGetQuestions(t *testing.T) {
	repo := repo.NewInMemoryRepo()
	domain := domain.NewQuizService(repo)
	server := NewRouter(domain)

	t.Run("Get Questions returns data from db", func(t *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, "/questions", nil)

		want := []model.Question{
			{ID: 1, Question: "What is FastTrack model based on?", Answers: []string{"Complexity", "Singularity", "Dangerous", "Fun"}},
			{ID: 2, Question: "When was FastTrack founded?", Answers: []string{"2010", "2013", "2016", "2019"}},
			{ID: 3, Question: "What is FastTrack core product?", Answers: []string{"iGaming CRM ", "Real state CRM", "Financial Services", "E-commerce"}},
			{ID: 4, Question: "What was last FastTrack component launch?", Answers: []string{"Singularity ", "Greco", "Vimeo", "Rewards"}},
		}
		server.ServeHTTP(response, request)
		got := extractFromResponse[[]model.Question](t, response.Body)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponse(t, got, want)
	})
}

func TestSubmitAnswers(t *testing.T) {
	repo := repo.NewInMemoryRepo()
	domain := domain.NewQuizService(repo)
	server := NewRouter(domain)

	t.Run("returns result with valid input valid", func(t *testing.T) {
		userAnswers := []model.Answer{
			{QuestionID: 1, UserAnswer: 0},
			{QuestionID: 2, UserAnswer: 1},
			{QuestionID: 3, UserAnswer: 2},
			{QuestionID: 4, UserAnswer: 3},
		}
		want := model.Result{CorrectAnswers: 1, Score: 13}

		body, _ := json.Marshal(userAnswers)
		request, _ := http.NewRequest(http.MethodPost, "/submit", bytes.NewReader([]byte(body)))
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		got := extractFromResponse[model.Result](t, response.Body)

		assertStatus(t, response.Code, http.StatusCreated)
		assertResponse(t, got, want)
	})
	t.Run("returns bad request when has more than one answer to the same question", func(t *testing.T) {
		userAnswers := []model.Answer{
			{QuestionID: 1, UserAnswer: 0},
			{QuestionID: 1, UserAnswer: 1},
			{QuestionID: 3, UserAnswer: 2},
			{QuestionID: 4, UserAnswer: 3},
		}

		body, _ := json.Marshal(userAnswers)
		request, _ := http.NewRequest(http.MethodPost, "/submit", bytes.NewReader([]byte(body)))
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusBadRequest)
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("response status is wrong, got %d want %d", got, want)
	}
}

func assertResponse[T any](t testing.TB, got, want T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("response body is wrong, got %v want %v", got, want)
	}
}

func extractFromResponse[T any](t testing.TB, body io.Reader) T {
	t.Helper()
	var got T
	err := json.NewDecoder(body).Decode(&got)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Question, '%v'", body, err)
	}
	return got
}
