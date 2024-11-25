package store

import (
	"sync"

	"github.com/lf-silva/fastTrack/model"
)

type InMemoryQuizStore struct {
	store []model.Question
	lock  sync.RWMutex
}

func NewInMemoryQuizStore() *InMemoryQuizStore {

	return &InMemoryQuizStore{
		store,
		sync.RWMutex{},
	}
}

func (q *InMemoryQuizStore) GetQuestions() []model.Question {
	return q.store
}

var store = []model.Question{
	{Question: "", Answers: []string{}, CorrectAnswer: 0},
	{Question: "", Answers: []string{}, CorrectAnswer: 0},
	{Question: "", Answers: []string{}, CorrectAnswer: 0},
}
