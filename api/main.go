package main

import (
	//"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type question struct {
	ID            int       `json:"id"`
	Question      string    `json:"question"`
	AllAnswers    [4]string `json:"allAnswers"`
	CorrectAnswer int       `json:"correctAnswer"`
}

type quizResult struct {
	QuizAnswers [4]int `json:"quizAnswers"`
}

var questions = []question{
	{ID: 1, Question: "Q1", AllAnswers: [4]string{"w1", "w1", "w1", "c1"}, CorrectAnswer: 4},
	{ID: 2, Question: "Q2", AllAnswers: [4]string{"w2", "c2", "w2", "w2"}, CorrectAnswer: 2},
	{ID: 3, Question: "Q3", AllAnswers: [4]string{"w3", "c3", "w3", "w3"}, CorrectAnswer: 2},
	{ID: 4, Question: "Q4", AllAnswers: [4]string{"c4", "w4", "w4", "w4"}, CorrectAnswer: 1},
}

var otherResults = []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

func getQuiz(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, questions)
}

func postResults(c *gin.Context) {
	var results quizResult
	var correctAnswers int = 0
	var correctPercentage int = 0
	if err := c.BindJSON(&results); err != nil {
		return
	}

	for i, q := range results.QuizAnswers {
		if questions[i].CorrectAnswer == q {
			correctAnswers++
		}
	}

	correctPercentage = correctAnswers * 100 / len(questions)
	fmt.Printf("You have answered %d questions correctly\n", correctAnswers)
	compareResult(correctPercentage)
}

func compareResult(result int) {
	var worseResults int = 0

	for _, i := range otherResults {
		if result > i {
			worseResults++
		}
	}

	var betterThan int = worseResults * 100 / len(otherResults)
	fmt.Printf("Your result is better than %d%% of the participants", betterThan)
}

func main() {
	router := gin.Default()
	router.GET("/quiz", getQuiz)
	router.Run("localhost:8080")
	router.POST("/quiz", postResults)
}
