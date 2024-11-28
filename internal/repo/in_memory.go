package repo

import (
	"sort"
	"sync"

	"github.com/google/uuid"
	"github.com/lf-silva/fastTrack/internal/model"
	"github.com/wangjia184/sortedset"
)

type InMemoryRepo struct {
	questions map[int]model.Question
	scores    *sortedset.SortedSet
	lock      sync.RWMutex
}

func NewInMemoryRepo() *InMemoryRepo {

	return &InMemoryRepo{
		storedQuestions,
		newSortedSet(),
		sync.RWMutex{},
	}
}

func (q *InMemoryRepo) GetQuestions() []model.Question {
	q.lock.RLock()
	defer q.lock.RUnlock()

	questions := make([]model.Question, 0, len(q.questions))
	for _, q := range q.questions {
		questions = append(questions, q)
	}
	sort.Slice(questions, func(i, j int) bool {
		return questions[i].ID < questions[j].ID
	})

	return questions
}

func (r *InMemoryRepo) GetQuestion(id int) (model.Question, bool) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	question, ok := r.questions[id]
	return question, ok
}

func (r *InMemoryRepo) SubmitScore(score int) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.scores.AddOrUpdate(uuid.NewString(), sortedset.SCORE(score), nil)
}

func (r *InMemoryRepo) GetIndexes(score int) (int, int) {
	r.lock.Lock()
	defer r.lock.Unlock()

	scoreRange := r.scores.GetByScoreRange(0, sortedset.SCORE(score), &sortedset.GetByScoreRangeOptions{
		ExcludeEnd: true,
	})
	return len(scoreRange), r.scores.GetCount()
}

var storedQuestions = map[int]model.Question{
	1: {ID: 1, Question: "What is FastTrack model based on?", Answers: []string{"Complexity", "Singularity", "Dangerous", "Fun"}, CorrectAnswer: 1},
	2: {ID: 2, Question: "When was FastTrack founded?", Answers: []string{"2010", "2013", "2016", "2019"}, CorrectAnswer: 2},
	3: {ID: 3, Question: "What is FastTrack core product?", Answers: []string{"iGaming CRM ", "Real state CRM", "Financial Services", "E-commerce"}, CorrectAnswer: 0},
	4: {ID: 4, Question: "What was last FastTrack component launch?", Answers: []string{"Singularity ", "Greco", "Vimeo", "Rewards"}, CorrectAnswer: 3},
}

func newSortedSet() *sortedset.SortedSet {
	s := sortedset.New()
	s.AddOrUpdate(uuid.New().String(), 0, nil)
	s.AddOrUpdate(uuid.New().String(), 1, nil)
	s.AddOrUpdate(uuid.New().String(), 4, nil)
	s.AddOrUpdate(uuid.New().String(), 2, nil)
	s.AddOrUpdate(uuid.New().String(), 3, nil)
	s.AddOrUpdate(uuid.New().String(), 4, nil)
	s.AddOrUpdate(uuid.New().String(), 3, nil)

	return s
}
