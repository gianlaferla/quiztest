package main

import (
	//"errors"
	"fmt"
	"net/http"

	models "QuizTest/models"

	"github.com/gin-gonic/gin"
)

var questions = []models.Question{
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
	var results models.Answers
	var correctAnswers int = 0
	var correctPercentage int = 0
	if err := c.BindJSON(&results); err != nil {
		return
	}

	for i, q := range results.Answers {
		if questions[i].CorrectAnswer == q {
			correctAnswers++
		} else {
			fmt.Printf("%d != %d", questions[i].CorrectAnswer, q)
		}
	}

	fmt.Printf("correct answers = %d ", correctAnswers)
	correctPercentage = correctAnswers * 100 / len(questions)

	comparison := compareResult(correctPercentage)
	toReturn := fmt.Sprintf("You have answered %d questions correctly. %s", correctAnswers, comparison)
	c.IndentedJSON(http.StatusAccepted, toReturn)
}

func compareResult(result int) string {
	var worseResults int = 0

	for _, i := range otherResults {
		if result > i {
			worseResults++
		}
	}

	var betterThan int = worseResults * 100 / len(otherResults)
	return fmt.Sprintf("Your result is better than %d%% of the participants", betterThan)
}

func main() {
	router := gin.Default()
	router.GET("/quiz", getQuiz)
	router.POST("/answers", postResults)
	router.Run("localhost:8080")
}
