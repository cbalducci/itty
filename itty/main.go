package main

import (
        "encoding/json"
        "net/http"
        "github.com/aws/aws-lambda-go/events"
        "github.com/aws/aws-lambda-go/lambda"
)

type SlackMessage struct {
	Response_Type  string  `json:"response_type"`
	Text           string  `json:"text"`
}

var terratest_repo = "https://github.com/gruntwork-io/terratest"
var answer string

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  resp, _ := http.Get(terratest_repo)
  if resp.StatusCode == 404 {
    answer = "Not yet, 404 repo not found"
    } else {
    answer = "Celebrate! It looks like the repo is accessible!"
  }

  json_answer, _ := json.Marshal(SlackMessage{Response_Type: "in_channel", Text: answer})
  return events.APIGatewayProxyResponse{Body: string(json_answer), StatusCode: 200}, nil
}

func main() {
  lambda.Start(Handler)
}
