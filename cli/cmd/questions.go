/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	models "QuizTest/models"

	"github.com/spf13/cobra"
)

// questionsCmd represents the questions command
var questionsCmd = &cobra.Command{
	Use:   "questions",
	Short: "This command is used to view all the questions",
	Long:  `This command is used to view all the questions in the quiz`,
	Run: func(cmd *cobra.Command, args []string) {
		getQuestions()
	},
}

func init() {
	rootCmd.AddCommand(questionsCmd)
}

func getQuestions() {
	url := "http://localhost:8080/quiz"
	responseBytes := getQuestionsData(url)
	allQuestions := []models.Question{}

	if err := json.Unmarshal(responseBytes, &allQuestions); err != nil {
		fmt.Printf("Could not unmarshal response")
		return
	}

	for i, question := range allQuestions {
		fmt.Printf("Question %d: %s\n\n", i+1, question.Question)
		for i2, answer := range question.AllAnswers {
			fmt.Printf("%d) %s\n", i2+1, answer)
		}
		fmt.Println("-----------------------------------------------------------------")
	}
}

func getQuestionsData(baseApi string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseApi,
		nil,
	)
	if err != nil {
		fmt.Printf("Could not create the request:\n\n%s\n", err)
	}

	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("Could not request the questions:\n\n%s\n", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Could not read the body:\n\n%s\n", err)
	}

	return responseBytes
}
