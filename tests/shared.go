package cmapv2tests

import (
	"crypto/rand"
)

type KeyVal struct {
	Key []byte
	Value []byte
}

func GenerateRandomBytes(length int) ([]byte, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil { return nil, err }
	return randomBytes, nil
}
