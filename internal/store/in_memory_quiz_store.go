package store

import (
	"sync"

	"github.com/lf-silva/fastTrack/internal/model"
)

type InMemoryQuizStore struct {
	store  map[int]model.Question
	scores []int
	lock   sync.RWMutex
}

func NewInMemoryQuizStore() *InMemoryQuizStore {

	return &InMemoryQuizStore{
		questions,
		scores,
		sync.RWMutex{},
	}
}

func (q *InMemoryQuizStore) GetQuestions() []model.Question {
	array := make([]model.Question, 0, len(questions))
	for _, q := range questions {
		array = append(array, q)
	}
	return array
}

func (q *InMemoryQuizStore) GetQuestion(id int) (model.Question, bool) {
	question, ok := questions[id]
	return question, ok
}

func (q *InMemoryQuizStore) SubmitScore(score int) {
	q.scores = append(q.scores, score)
}

var questions = map[int]model.Question{
	1: {ID: 1, Question: "What is FastTrack model based on?", Answers: []string{"Complexity", "Singularity", "Dangerous", "Fun"}, CorrectAnswer: 1},
	2: {ID: 2, Question: "When was FastTrack founded?", Answers: []string{"2010", "2013", "2016", "2019"}, CorrectAnswer: 2},
	3: {ID: 3, Question: "What is FastTrack core product?", Answers: []string{"iGaming CRM ", "Real state CRM", "Financial Services", "E-commerce"}, CorrectAnswer: 0},
	4: {ID: 4, Question: "What was last FastTrack component launch?", Answers: []string{"Singularity ", "Greco", "Vimeo", "Rewards"}, CorrectAnswer: 3},
}

var scores = []int{3, 4, 0, 1, 2}
