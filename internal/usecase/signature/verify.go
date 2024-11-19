package signature

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"os"
)

type VerifyInterface interface {
	Handle(headers map[string]string, body string) bool
}

func Handle(headers map[string]string, body string) bool {
	signature := headers["x-signature-ed25519"]
	timestamp := headers["x-signature-timestamp"]

	fmt.Println("Signature: ", signature)
	fmt.Println("Timestamp: ", timestamp)

	// Verify the request signature
	isVerified := verifySignature(signature, timestamp, body)
	fmt.Println("isVerified: ", isVerified)

	return isVerified
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
