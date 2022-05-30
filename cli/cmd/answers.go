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

	"github.com/spf13/cobra"
)

// answersCmd represents the answers command
var answersCmd = &cobra.Command{
	Use:   "answers",
	Short: "Pass the answers and get the results",
	Long:  `Pass the answers and get the results`,
	Run: func(cmd *cobra.Command, args []string) {
		postAnswers()
	},
}

func postAnswers() {
	url := "http://localhost:8080/answers"
	responseBytes := postAnswersData(url)
	fmt.Println(string(responseBytes))
}

func postAnswersData(baseApi string) []byte {
	answerData := quizstructs.answers{Answers: [4]int{4, 2, 2, 1}}
	json_data, err := json.Marshal(answerData)

	fmt.Println(string(json_data))

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

func init() {
	rootCmd.AddCommand(answersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// answersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// answersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
