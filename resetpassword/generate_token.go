package resetpassword

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateResetToken() (string, error) {
	token := make([]byte, 32) // Generate a 32-byte random token
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(token), nil
}

// func StoreResetToken(token string) error{

// }
