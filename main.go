package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Message string `json:"message"`
}

func handler(ctx context.Context) (Response, error) {
	fmt.Println(ctx)
	fmt.Println("Hello, World!")

	resp := Response{
		Message: "Olá, mundo!",
	}

	return resp, nil
}

func main() {
	lambda.Start(handler)
}
