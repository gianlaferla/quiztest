/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	models "QuizTest/models"

	"github.com/spf13/cobra"
)

var answersCmd = &cobra.Command{
	Use:   "answers",
	Short: "Post answers to quiz api",
	Long:  `Post the answers from the questions.json file to the quiz api to get the results.`,
	Run: func(cmd *cobra.Command, args []string) {
		postAnswers()
	},
}

var answers = models.Answers{}

func postAnswers() {
	url := "http://localhost:8080/answers"
	responseBytes := postAnswersData(url)
	fmt.Println(string(responseBytes))
}

func postAnswersData(baseApi string) []byte {
	initAnswers()
	json_data, err := json.Marshal(answers)

	if err != nil {
		fmt.Println(err)
	}

	response, err := http.Post(
		baseApi,
		"application/json",
		bytes.NewBuffer(json_data),
	)

	if err != nil {
		fmt.Printf("Could not create the POST request:\n\n%s\n", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Could not read the body:\n\n%s\n", err)
	}

	return responseBytes
}

func initAnswers() {
	answersJSONFile := "answers.json"
	file, err := ioutil.ReadFile("../" + answersJSONFile)

	if err != nil {
		fmt.Printf("Unable to read %s", answersJSONFile)
		os.Exit(4)
	} else {
		err = json.Unmarshal([]byte(file), &answers)
	}

	if err != nil {
		fmt.Printf("Unable to unmarshal data from %s", answersJSONFile)
		os.Exit(4)
	}
}

func init() {
	rootCmd.AddCommand(answersCmd)
}
