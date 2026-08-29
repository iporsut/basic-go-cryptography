//go:build ignore

package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"log"
)

func main() {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	message := []byte("example message")
	signature := ed25519.Sign(privateKey, message)
	// signature = signature[:len(signature)-1] // remove last byte to simulate tampering
	valid := ed25519.Verify(publicKey, message, signature)
	if !valid {
		log.Fatal("signature verification failed")
	}
}
