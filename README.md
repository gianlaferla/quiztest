# quiztest
This project in Golang is made up of 2 components, a quiz api, and a cli tool which interacts with the api.

The 'api' folder contains a simple api which accepts the following requests:
  1) GET request which returns all the questions of the quiz.
  2) POST request with the answers attached in the json, returning the result. 

How to run

1) First run the api server locally in cmd. The questions for the api are stored in 'answers.json' which should be in the relative path. Run the following: 
  - cd quiztest/api 
  - go run main.go

2) Build the cli.exe found in the cli folder. Run the following
  - cd quiztest/cli
  - go build

3) Once the cli.exe is created, this may be used together with the 'questions' parameter, to display a set of questions. Run the following:
  - cd quiztest/cli
  - cli questions

4) The parameter 'answers' may be used to post the answers stored in answers.json to the api. A result will be shown in the cmd. Run the following:
  - cd quiztest/cli
  - cli answers
