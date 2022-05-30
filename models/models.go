package models

type Question struct {
	ID            int       `json:"id"`
	Question      string    `json:"question"`
	AllAnswers    [4]string `json:"allAnswers"`
	CorrectAnswer int       `json:"correctAnswer"`
}

type Answers struct {
	Answers [4]int `json:"Answers"`
}
