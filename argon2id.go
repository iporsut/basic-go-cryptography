//go:build ignore

package main

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/argon2"
)

func main() {
	salt, _ := hex.DecodeString("59d7bf2721948c9c55d2c169db124c7f")
	fmt.Println("salt:", hex.EncodeToString(salt))

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
