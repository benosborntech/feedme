package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRand(length int) (string, error) {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	state := base64.URLEncoding.EncodeToString(bytes)

	return state, nil
}
