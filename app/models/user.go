package models

import (
	"crypto/rand"
	"encoding/hex"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateSecureToken() string {
	b := make([]byte, 20)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
