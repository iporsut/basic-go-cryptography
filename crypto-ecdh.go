//go:build ignore

package main

import (
	"bytes"
	"crypto/ecdh"
	"crypto/hkdf"
	"crypto/rand"
	"crypto/sha256"
	"log"
)

func main() {
	curve := ecdh.X25519()
	alice, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	alicePub := alice.PublicKey()

	bob, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	bobPub := bob.PublicKey()

	// Derive shared secret
	aliceShared, err := alice.ECDH(bobPub)
	if err != nil {
		log.Fatal(err)
	}

	bobShared, err := bob.ECDH(alicePub)
	if err != nil {
		log.Fatal(err)
	}

	if !bytes.Equal(aliceShared, bobShared) {
		log.Fatal("shared secrets do not match")
	}

	salt := []byte("example salt")
	info := "example info"

	// derive a session key from the shared secret
	aliceSessionKey, err := hkdf.Key(sha256.New, aliceShared, salt, info, 32)
	if err != nil {
		log.Fatal(err)
	}

	bobSessionKey, err := hkdf.Key(sha256.New, bobShared, salt, info, 32)
	if err != nil {
		log.Fatal(err)
	}

	if !bytes.Equal(aliceSessionKey, bobSessionKey) {
		log.Fatal("session keys do not match")
	}
}
