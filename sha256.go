//go:build ignore

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	// exampel sha256
	data := []byte("example data")
	hash := sha256.Sum256(data)
	fmt.Println("SHA256:", hex.EncodeToString(hash[:]))
}
