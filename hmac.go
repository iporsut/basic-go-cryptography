//go:build ignore

package main

import (
	"crypto/hmac"
	"crypto/sha256"
)

func main() {
	key := []byte("secret key")
	mac := hmac.New(sha256.New, key)
	data := []byte("example data")
	mac.Write(data)
	signature := mac.Sum(nil)

	// verify the signature
	expectedMac := hmac.New(sha256.New, key)
	expectedMac.Write(data)
	expectedSignature := expectedMac.Sum(nil)

	if hmac.Equal(signature, expectedSignature) {
		println("Signature is valid")
	} else {
		println("Signature is invalid")
	}
}
