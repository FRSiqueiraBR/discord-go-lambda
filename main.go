package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/FRSiqueiraBR/discord-go-lambda/internal/usecase/signature"
)

type BodyResponse struct {
	Type int `json:"type"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	isVerified := signature.Handle(request.Headers, request.Body)

	if !isVerified {
		return events.APIGatewayProxyResponse{
			Body:       "invalid request signature",
			StatusCode: 401,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       "{\"type\": 4, \"data\": {\"content\": \"Hello, World.\"}}",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
