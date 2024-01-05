package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type BodyResponse struct {
	Type int `json:"type"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	rawBody := request.Body
	fmt.Println("Start")
	fmt.Println(request.Body)
	fmt.Println(request.Headers)
	fmt.Println(request)
	fmt.Println("Body: ", rawBody)

	return events.APIGatewayProxyResponse{
		Body:       "{\"type\": 1}",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
