package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Nyxoy/restAPI/db"
	"github.com/Nyxoy/restAPI/models"
	"github.com/Nyxoy/restAPI/resetpassword"
	"github.com/Nyxoy/restAPI/utils"
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

func HandleForgot(w http.ResponseWriter, r *http.Request) {
	var fuser models.Forgot
	var users []models.User
	// Validating
	validate := utils.NewValidator()
	if err := validate.Struct(fuser); err != nil {
		//Gather the error data
		errors := err.(validator.ValidationErrors)
		for _, value := range errors {
			log.Printf("Validation failed for field %s , conditon %s /n", value.Field(), value.Tag())
		}
		utils.WriteError(w, http.StatusBadRequest, "Invalid Input Fields")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&fuser); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid body")
		return
	}

	db := db.CreateRestyClient()
	resp, err := db.R().
		SetQueryParam("email", "eq."+fuser.Email). // Set the query parameter for email
		Get(viper.GetString("DB_BASE_URL") + "/rest/v1/users")
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "An error occured")
		return
	}
	var exists bool = false
	if resp.StatusCode() == 200 {

		if err := json.Unmarshal(resp.Body(), &users); err != nil {
			utils.WriteError(w, http.StatusBadRequest, "Wrong data fetched")
			return
		}

		if len(users) > 0 {
			exists = true
		}
	}
	if exists {
		token, err := resetpassword.GenerateResetToken()
		if err != nil {
			log.Println("Error occured at the time of token generation")
			return
		}
		// Store the data in the reset_password table
		actUser := users[0]
		expire := time.Now().Add(30 * time.Second)
		data := map[string]interface{}{
			"user_id":     actUser.User_ID,
			"reset_token": token,
			"expiration":  expire,
			"used":        false,
		}
		resp1, err1 := db.R().SetBody(data).Post(viper.GetString("DB_BASE_URL") + "/rest/v1/password_reset_tokens")
		if err1 != nil {
			utils.WriteError(w, http.StatusInternalServerError, "Unable to register the token")
			log.Printf("An error happend when registering the token for %s %s", fuser.Email, err1)
			return
		}
		fmt.Println(resp1.StatusCode())
		if resp1.StatusCode() == 201 {
			// utils.WriteError(w, http.StatusOK, "If your mail is registerd a mail will be sent on the mail id")
			mail := EmailBody{
				Toaddr:  actUser.Email,
				Subject: "Password reset Link",
				Body:    "Please use the token in the url for verification and reseting the password" + token,
			}
			HandleEmail(w, r, mail)
		}
		// utils.WriteError(w, http.StatusOK, "If your mail is registerd a mail will be sent on the mail id")
		// Sending the mail logic here

	} else {
		utils.WriteError(w, http.StatusOK, "If your mail is registerd a mail will be sent on the mail id")
	}

}
