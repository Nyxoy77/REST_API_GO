package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Nyxoy/restAPI/models"
	"github.com/Nyxoy/restAPI/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateToken(user models.User) (string, error) {
	secret_key := viper.GetString("SECRET_KEY")
	claims := jwt.MapClaims{
		"user_id":   user.User_ID,
		"email":     user.Email,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
		"issued_at": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret_key))
}

func VerifyJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			utils.WriteError(w, http.StatusUnauthorized, "Missing Authorization header")
			return
		}
		tokenString = tokenString[len("Bearer "):]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(viper.GetString("SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			utils.WriteError(w, http.StatusUnauthorized, "Invalid Token")
			return
		}
		next(w, r)
	}
}
