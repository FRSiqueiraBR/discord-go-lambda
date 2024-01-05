package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	rawBody := request.Body
	fmt.Println("Start")
	fmt.Println("Body: ", rawBody)

	json, _ := json.Marshal(
		struct {
			Type int `json:"type"`
		}{
			Type: 1,
		},
	)

	return events.APIGatewayProxyResponse{
		Body:       string(json),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
