//go:build ignore

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	privateKey, err := rsa.GenerateKey(
		rand.Reader,
		3072,
	)

	if err != nil {
		log.Fatal(err)
	}

	publicKey := &privateKey.PublicKey
	label := []byte("OAEP Encrypted")

	ciphertext, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		[]byte("example plaintext"),
		label,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(ciphertext))

	decrypted, err := rsa.DecryptOAEP(
		sha256.New(),
		nil,
		privateKey,
		ciphertext,
		label,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Decrypted:", string(decrypted))
}
