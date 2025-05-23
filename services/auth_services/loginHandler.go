package authservices

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Nyxoy/restAPI/db"
	"github.com/Nyxoy/restAPI/models"
	myjwt "github.com/Nyxoy/restAPI/services/jwt_logic"
	"github.com/Nyxoy/restAPI/utils"
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.Login

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return
	}
	validate := utils.NewValidator()
	if err := validate.Struct(user); err != nil {
		//Gather the error data
		errors := err.(validator.ValidationErrors)
		for _, value := range errors {
			log.Printf("Validation failed for field %s , conditon %s /n", value.Field(), value.Tag())
		}
		utils.WriteError(w, http.StatusBadRequest, "Invalid Input Fields")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	db := db.CreateRestyClient()
	resp, err := db.R().SetContext(ctx).
		SetQueryParam("email", "eq."+user.Email).
		Get(viper.GetString("DB_BASE_URL") + "/rest/v1/users")
	if err != nil {
		log.Printf("An error occured at the time of logging %v", err)
		http.Error(w, fmt.Sprintf("An error occurred at the time of logging: %s", err), http.StatusInternalServerError)
		return
	}
	var respUser []models.User
	err1 := json.Unmarshal(resp.Body(), &respUser)
	if err1 != nil {
		log.Printf("An error occured at the time of parsing the data %v", err)
		http.Error(w, "An error occured at the time of parsing the data for logging", http.StatusInternalServerError)
		return
	}
	if len(respUser) == 0 {
		utils.WriteError(w, http.StatusUnauthorized, "User not found please register")
		return
	}

	var actUser = respUser[0]

	matches := utils.CheckHashPass(user.Password, actUser.Password)

	if matches {
		token, err1 := myjwt.GenerateToken(actUser.User_ID, actUser.Email, actUser.UserType)
		if err1 != nil {
			utils.WriteError(w, http.StatusInternalServerError, "An error occured during token generation")
			return
		}
		refreshToken, err2 := myjwt.GenerateRefreshToken(actUser.User_ID, actUser.Email, actUser.UserType)
		if err2 != nil {
			utils.WriteError(w, http.StatusInternalServerError, "An error occured during refresh token generation")
			return
		}
		data := map[string]interface{}{
			"message":       "The user is successfully logged in",
			"email":         actUser.Email,
			"status":        http.StatusOK,
			"token":         token,
			"refresh_token": refreshToken,
		}
		json.NewEncoder(w).Encode(data)
		// utils.WriteError(w, http.StatusFound, "The user successfully logged in")
	} else {
		utils.WriteError(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}
}
