package model

type Result struct {
	CorrectAnswers int     `json:"correctAnswers"`
	Score          float64 `json:"score"`
}
