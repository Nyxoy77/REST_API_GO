package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Nyxoy/restAPI/models"
	"github.com/Nyxoy/restAPI/utils"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func GenerateToken(user_id int, user_email string, role string) (string, error) {
	secret_key := viper.GetString("SECRET_KEY")
	claims := &models.Claims{
		User_ID:  user_id,
		Email:    user_email,
		UserType: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret_key))
}
func GenerateRefreshToken(user_id int, user_email string, role string) (string, error) {
	secret_key := viper.GetString("SECRET_KEY")
	claims := &models.Claims{
		User_ID:  user_id,
		Email:    user_email,
		UserType: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
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
		// Parse only if you want to check the signature if it matches
		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(viper.GetString("SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			utils.WriteError(w, http.StatusUnauthorized, "Invalid Token")
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			utils.WriteError(w, http.StatusUnauthorized, "Token has expired")
			return
		}

		// IF the token has been expired
		// Automatically hojata hai check jwt.ParseWithClaims se
		// Setting the claims in the request context
		r = r.WithContext(context.WithValue(r.Context(), "claims", claims))
		next(w, r)
	}
}

type Refresh struct {
	Token string `json:"token"` // Use a more appropriate key name
}

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	var ref Refresh
	if err := json.NewDecoder(r.Body).Decode(&ref); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid body")
		return
	}

	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(ref.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("SECRET_KEY")), nil
	})

	// Check for token validity and expiration
	if err != nil || !token.Valid || claims.ExpiresAt < time.Now().Unix() {
		utils.WriteError(w, http.StatusUnauthorized, "Invalid or expired refresh token. Please login again.")
		return
	}

	// Generate a new access token (and refresh token if needed)
	newAccessToken, err := GenerateToken(claims.User_ID, claims.Email, claims.UserType)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error generating new access token")
		return
	}

	// Return the new access token in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token": newAccessToken,
	})
}
