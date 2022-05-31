package main

import (
	//"errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	models "QuizTest/models"

	"github.com/gin-gonic/gin"
)

var otherResults = []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
var questions = []models.Question{}

func initQuestions() {
	questionsJsonFileName := "questions.json"
	file, err := ioutil.ReadFile("../" + questionsJsonFileName)

	if err != nil {
		fmt.Printf("Unable to read %s file\n", questionsJsonFileName)
	} else {
		err = json.Unmarshal([]byte(file), &questions)
	}

	if err != nil {
		fmt.Printf("Unable to read the %s file\n", questionsJsonFileName)
		return
	}
}

func getQuiz(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, questions)
}

func postResults(c *gin.Context) {
	results := models.Answers{}
	correctAnswers := 0
	correctPercentage := 0
	if err := c.BindJSON(&results); err != nil {
		return
	}

	for i, q := range results.Answers {
		if i < len(questions) && questions[i].CorrectAnswer == q {
			correctAnswers++
		}
	}

	fmt.Printf("Correct answers = %d ", correctAnswers)
	correctPercentage = correctAnswers * 100 / len(questions)

	comparison := compareResult(correctPercentage)
	warning := ""

	wereInResponse := "were"
	answersInResponse := "answers"
	if len(results.Answers) == 1 {
		wereInResponse = "was"
		answersInResponse = "answer"
	}

	correctQuestionsInResponse := "questions"
	if correctAnswers == 1 {
		correctQuestionsInResponse = "question"
	}

	if len(questions) != len(results.Answers) {
		if len(questions) < len(results.Answers) {
			warning = fmt.Sprintf("Found %d %s. Only the first %d %s will be used. ", len(results.Answers), answersInResponse, len(questions), answersInResponse)
		} else {
			warning = fmt.Sprintf("Only %d %s %s found for %d questions. ", len(results.Answers), answersInResponse, wereInResponse, len(questions))
		}
	}
	toReturn := fmt.Sprintf("You have answered %d %s correctly. %s", correctAnswers, correctQuestionsInResponse, comparison)

	c.IndentedJSON(http.StatusAccepted, warning+toReturn)
}

func compareResult(result int) string {
	worseResults := 0

	for _, i := range otherResults {
		if result > i {
			worseResults++
		}
	}

	betterThan := worseResults * 100 / len(otherResults)
	return fmt.Sprintf("Your result is better than %d%% of the participants!", betterThan)
}

func main() {
	initQuestions()

	router := gin.Default()
	router.GET("/quiz", getQuiz)
	router.POST("/answers", postResults)
	router.Run("localhost:8080")
}
