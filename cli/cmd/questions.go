/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// questionsCmd represents the questions command
var questionsCmd = &cobra.Command{
	Use:   "questions",
	Short: "This command is used to view all the questions in the quiz",
	Long:  `This command is used to view all the questions in the quiz`,
	Run: func(cmd *cobra.Command, args []string) {
		getQuestions()
	},
}

func init() {
	rootCmd.AddCommand(questionsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// questionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// questionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type question struct {
	ID            int       `json:"id"`
	Question      string    `json:"question"`
	AllAnswers    [4]string `json:"allAnswers"`
	CorrectAnswer int       `json:"correctAnswer"`
}

func getQuestions() {
	url := "localhost:8080/quiz"
	responseBytes := getQuestionsData(url)
	allQuestions := []question{}

	if err := json.Unmarshal(responseBytes, &allQuestions); err != nil {
		fmt.Printf("Could not unmarshal response - %w", err)
	}

	for _, question := range allQuestions {
		fmt.Println(question.Question)
		for i, answer := range question.AllAnswers {
			fmt.Printf("%d) %s\n", i, answer)
		}
		fmt.Println("------------------------------------")
	}
}

func getQuestionsData(baseApi string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseApi,
		nil,
	)
	if err != nil {
		fmt.Printf("Could not request the questions - %w", err)
	}

	request.Header.Add("Accept", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("Could not request the questions - %w", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Could not read body - %w", err)
	}

	return responseBytes
}
