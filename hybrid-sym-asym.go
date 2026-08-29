//go:build ignore

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	// generate a random session encryption key
	sessionKey := make([]byte, 32)
	if _, err := rand.Read(sessionKey); err != nil {
		log.Fatal(err)
	}

	// encrypt the plaintext using the session key
	plaintext := []byte("example plaintext")

	// new cipher block
	block, err := aes.NewCipher(sessionKey)
	if err != nil {
		log.Fatal(err)
	}

	// new GCM mode AEAD cipher
	aead, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}

	// generate a random nonce
	nonce := make([]byte, aead.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		log.Fatal(err)
	}

	// additional authenticated data
	aad := []byte("example aad")

	seal := aead.Seal(nonce, nonce, plaintext, aad)

	fmt.Println("Encrypted data:", hex.EncodeToString(seal))

	// Encrypt session key using RSA public key
	privateKey, err := rsa.GenerateKey(
		rand.Reader,
		3072,
	)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := &privateKey.PublicKey
	label := []byte("OAEP Encrypted")

	encryptedSessionKey, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		sessionKey,
		label,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encrypted session key:", hex.EncodeToString(encryptedSessionKey))

	decryptedSessionKey, err := rsa.DecryptOAEP(
		sha256.New(),
		nil,
		privateKey,
		encryptedSessionKey,
		label,
	)

	if err != nil {
		log.Fatal(err)
	}

	newBlock, err := aes.NewCipher(decryptedSessionKey)
	if err != nil {
		log.Fatal(err)
	}

	newAead, err := cipher.NewGCM(newBlock)
	if err != nil {
		log.Fatal(err)
	}

	decrypted, err := newAead.Open(nil, seal[:newAead.NonceSize()], seal[newAead.NonceSize():], aad)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decrypted data:", string(decrypted))
}
