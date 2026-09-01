//go:build ignore

package main

import (
	"crypto/hpke"
	"log"
)

func main() {
	kem := hpke.MLKEM768X25519()
	kdf := hpke.HKDFSHA256()
	aead := hpke.AES256GCM()

	k, err := kem.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	recipientPrivateKey := k
	publicKeyBytes := k.PublicKey().Bytes()

	var ciphertext []byte

	// Sender side
	{
		publicKey, err := kem.NewPublicKey(publicKeyBytes)
		if err != nil {
			log.Fatal(err)
		}

		message := []byte("|-()-|")
		ct, err := hpke.Seal(publicKey, kdf, aead, []byte("example"), message)
		if err != nil {
			log.Fatal(err)
		}
		ciphertext = ct
	}

	// Recipient side
	{
		plaintext, err := hpke.Open(recipientPrivateKey, kdf, aead, []byte("example"), ciphertext)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Decrypted message: %s", string(plaintext))
	}
}
