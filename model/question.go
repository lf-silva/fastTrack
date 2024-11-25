package model

type Question struct {
	ID            int
	Question      string
	Answers       []string
	CorrectAnswer int
}
