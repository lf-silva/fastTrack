package domain_test

import (
	"reflect"
	"testing"

	"github.com/lf-silva/fastTrack/internal/domain"
	"github.com/lf-silva/fastTrack/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestGetQuestions(t *testing.T) {
	mockRepo := new(domain.MockQuizRepo)

	t.Run("returns questions", func(t *testing.T) {
		questions := []model.Question{
			{ID: 1, Question: "What is FastTrack model based on?", Answers: []string{"Complexity", "Singularity", "Dangerous", "Fun"}, CorrectAnswer: 1},
			{ID: 2, Question: "When was FastTrack founded?", Answers: []string{"2010", "2013", "2016", "2019"}, CorrectAnswer: 2},
			{ID: 3, Question: "What is FastTrack core product?", Answers: []string{"iGaming CRM ", "Real state CRM", "Financial Services", "E-commerce"}, CorrectAnswer: 0},
			{ID: 4, Question: "What was last FastTrack component launch?", Answers: []string{"Singularity ", "Greco", "Vimeo", "Rewards"}, CorrectAnswer: 3},
		}

		mockRepo.On("GetQuestions").Return(questions)

		domain := domain.NewQuizService(mockRepo)
		got := domain.GetQuestions()

		mockRepo.AssertNumberOfCalls(t, "GetQuestions", 1)
		assert.Equal(t, questions, got)
		assert.Len(t, mockRepo.Calls, 1)
	})
}

func TestSubmitAnswers(t *testing.T) {
	mockRepo := new(domain.MockQuizRepo)

	t.Run("validates questions and returns correct value", func(t *testing.T) {
		answers := []model.Answer{
			{QuestionID: 1, UserAnswer: 0},
			{QuestionID: 2, UserAnswer: 2},
		}
		questions := []model.Question{
			{ID: 1, Question: "What is FastTrack model based on?", Answers: []string{"Complexity", "Singularity", "Dangerous", "Fun"}, CorrectAnswer: 1},
			{ID: 2, Question: "When was FastTrack founded?", Answers: []string{"2010", "2013", "2016", "2019"}, CorrectAnswer: 2},
		}

		call1 := mockRepo.On("GetQuestion").Return(questions[0], true).Once()
		call2 := mockRepo.On("GetQuestion").Return(questions[1], true).NotBefore(call1)
		call3 := mockRepo.On("SubmitScore").Return().NotBefore(call1, call2)
		mockRepo.On("GetIndexes").Return(1, 10).NotBefore(call1, call2, call3)

		domain := domain.NewQuizService(mockRepo)

		want := model.Result{CorrectAnswers: 1, Score: 10}
		got, err := domain.SubmitAnswers(answers)

		mockRepo.AssertNumberOfCalls(t, "GetQuestion", 2)
		mockRepo.AssertNumberOfCalls(t, "SubmitScore", 1)
		mockRepo.AssertNumberOfCalls(t, "GetIndexes", 1)
		assert.NoError(t, err)
		assertResult(t, got, want)
	})

	t.Run("returns error when has more than one answer for the same question", func(t *testing.T) {
		mockRepo := new(domain.MockQuizRepo)

		d := domain.NewQuizService(mockRepo)

		answers := []model.Answer{
			{QuestionID: 1, UserAnswer: 1},
			{QuestionID: 1, UserAnswer: 2},
			{QuestionID: 3, UserAnswer: 0},
			{QuestionID: 4, UserAnswer: 3},
		}
		_, err := d.SubmitAnswers(answers)

		mockRepo.AssertNumberOfCalls(t, "GetQuestion", 0)
		mockRepo.AssertNumberOfCalls(t, "SubmitScore", 0)
		mockRepo.AssertNumberOfCalls(t, "GetIndexes", 0)
		assert.Error(t, err, domain.MoreThanOneAnswerProvided)
	})
}

func assertResult(t *testing.T, got, want model.Result) {
	t.Helper()
	// const equalityThreshold = 1e-7

	if !reflect.DeepEqual(got, want) {
		//got.CorrectAnswers != want.CorrectAnswers || math.Abs(got.Score-want.Score) > equalityThreshold {
		t.Errorf("result is wrong, got %v want %v", got, want)
	}
}
