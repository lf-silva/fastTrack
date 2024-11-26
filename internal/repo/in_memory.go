package repo

import (
	"sort"
	"sync"

	"github.com/lf-silva/fastTrack/internal/model"
)

type InMemoryRepo struct {
	store  map[int]model.Question
	scores []int
	lock   sync.RWMutex
}

func NewInMemoryRepo() *InMemoryRepo {

	return &InMemoryRepo{
		storedQuestions,
		scores,
		sync.RWMutex{},
	}
}

func (q *InMemoryRepo) GetQuestions() []model.Question {
	questions := make([]model.Question, 0, len(storedQuestions))
	for _, q := range storedQuestions {
		questions = append(questions, q)
	}
	sort.Slice(questions, func(i, j int) bool {
		return questions[i].ID < questions[j].ID
	})

	return questions
}

func (q *InMemoryRepo) GetQuestion(id int) (model.Question, bool) {
	question, ok := storedQuestions[id]
	return question, ok
}

func (q *InMemoryRepo) SubmitScore(score int) {
	q.scores = append(q.scores, score)
}

var storedQuestions = map[int]model.Question{
	1: {ID: 1, Question: "What is FastTrack model based on?", Answers: []string{"Complexity", "Singularity", "Dangerous", "Fun"}, CorrectAnswer: 1},
	2: {ID: 2, Question: "When was FastTrack founded?", Answers: []string{"2010", "2013", "2016", "2019"}, CorrectAnswer: 2},
	3: {ID: 3, Question: "What is FastTrack core product?", Answers: []string{"iGaming CRM ", "Real state CRM", "Financial Services", "E-commerce"}, CorrectAnswer: 0},
	4: {ID: 4, Question: "What was last FastTrack component launch?", Answers: []string{"Singularity ", "Greco", "Vimeo", "Rewards"}, CorrectAnswer: 3},
}

var scores = []int{3, 4, 0, 1, 2}
