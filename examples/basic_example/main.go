package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/snark87/secretguard"
)

func main() {
	// make sure that in case of panic we close the memguard session
	defer secretguard.EnsureSafePanic()

	secret := secretguard.NewSecretFromBytes([]byte("this string contains secret"))
	var hash [sha256.Size]byte

	secret.Use(func(rawSecret []byte) error {
		// Using rawSecret outside of closure results in undefined behavior

		hash = sha256.Sum256(rawSecret)
		return nil
	})

	fmt.Printf("hash of secret value: %x\n", hash)
}
