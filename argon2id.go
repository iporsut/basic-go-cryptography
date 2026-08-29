// go:build ignore
package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/argon2"
)

func main() {
	salt := make([]byte, 16)
	rand.Read(salt)

	password := []byte("example password")

	digest := argon2.IDKey(
		password,
		salt,
		3,       // time
		32*1024, // memory
		4,       // threads
		32,      // key length
	)

	fmt.Println("Argon2id digest:", hex.EncodeToString(digest))
}
