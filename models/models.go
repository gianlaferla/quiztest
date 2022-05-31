package models

type Question struct {
	ID            int      `json:"id"`
	Question      string   `json:"question"`
	AllAnswers    []string `json:"allAnswers"`
	CorrectAnswer int      `json:"correctAnswer"`
}

type Answers struct {
	Answers []int `json:"Answers"`
}
