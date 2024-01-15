package main

import (
	"fmt"
	"os"

	"crypto/ed25519"
	"encoding/hex"

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

	signature := request.Headers["X-Signature-Ed25519"]
	timestamp := request.Headers["X-Signature-Timestamp"]

	fmt.Println("Signature: ", signature)
	fmt.Println("Timestamp: ", timestamp)

	// Verify the request signature
	isVerified := verifySignature(signature, timestamp, rawBody)
	fmt.Println("isVerified: ", isVerified)

	// Respond accordingly
	if !isVerified {
		return events.APIGatewayProxyResponse{
			Body:       "invalid request signature",
			StatusCode: 401,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       "{\"type\": 1}",
		StatusCode: 200,
	}, nil
}

func verifySignature(signature string, timestamp string, body string) bool {
	// Concatenate timestamp and body
	message := timestamp + body

	// Get application public key
	applicationPublicKey := os.Getenv("APPLICATION_PUBLIC_KEY")
	fmt.Println("APPLICATION_PUBLIC_KEY: ", applicationPublicKey)

	// Decode hex signature and public key
	decodedSignature, _ := hex.DecodeString(signature)
	decodedPublicKey, _ := hex.DecodeString(applicationPublicKey)

	fmt.Println("Signature: ", decodedSignature)
	fmt.Println("PublicKey: ", decodedPublicKey)

	// Verify the signature
	return ed25519.Verify(decodedPublicKey, []byte(message), decodedSignature)
}

func main() {
	lambda.Start(handler)
}
