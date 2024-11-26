package model

type Question struct {
	ID            int      `json:"id"`
	Question      string   `json:"question"`
	Answers       []string `json:"answers"`
	CorrectAnswer int      `json:"-"`
}

func (q *Question) IsAnswerCorrect(userAnswer int) bool {
	return q.CorrectAnswer == userAnswer
}
