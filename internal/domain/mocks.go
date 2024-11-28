package domain

import (
	"github.com/lf-silva/fastTrack/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockQuizRepo struct {
	mock.Mock
}

func (mock *MockQuizRepo) GetQuestions() []model.Question {
	args := mock.Called()

	return args.Get(0).([]model.Question)
}

func (mock *MockQuizRepo) GetQuestion(id int) (model.Question, bool) {
	args := mock.Called()

	return args.Get(0).(model.Question), args.Bool(1)
}

func (mock *MockQuizRepo) SubmitScore(result int) {
	mock.Called()
}

func (mock *MockQuizRepo) GetIndexes(score int) (int, int) {
	args := mock.Called()

	return args.Int(0), args.Int(1)
}
