package services

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomToken() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
