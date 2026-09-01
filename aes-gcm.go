//go:build ignore

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	// make random key
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	// new cipher block
	block, err := aes.NewCipher(key)
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

	plaintext := []byte("example plaintext")

	// additional authenticated data
	aad := []byte("example aad")

	sealed := aead.Seal(nonce, nonce, plaintext, aad)

	fmt.Println("sealed:", hex.EncodeToString(sealed))
	fmt.Println("nonce:", hex.EncodeToString(nonce))

	// decrypt the sealed ciphertext
	// we need to use same key, nonce, and aad to decrypt
	nonce = sealed[:aead.NonceSize()]
	sealed = sealed[aead.NonceSize():]
	decrypted, err := aead.Open(nil, nonce, sealed, aad)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(decrypted))
}
