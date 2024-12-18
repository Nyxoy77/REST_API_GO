package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alexedwards/argon2id"
)

func ParseJSON(r *http.Request, payload any, w http.ResponseWriter) error {
	if r.Body == nil {
		w.Write([]byte("The body is nill please enter the body parameters"))
		return fmt.Errorf("the body is empty cant be empty")
	}
	return json.NewDecoder(r.Body).Decode(&payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, mess string) {
	WriteJSON(w, status, map[string]interface{}{
		"message":    mess,
		"statusCode": status,
	})
}

func Encrypt(pass string) (string, error) {
	hash, err := argon2id.CreateHash(pass, argon2id.DefaultParams)
	if err != nil {
		log.Println("There was error hasing the password")
		return "", err
	}
	return hash, nil
}

func CheckHashPass(plainpass, hashPass string) bool {
	match, err := argon2id.ComparePasswordAndHash(plainpass, hashPass)
	if err != nil {
		log.Println("Error occured while comparing the passwords")
		return false
	}
	return match
}
