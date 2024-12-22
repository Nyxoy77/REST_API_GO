package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Nyxoy/restAPI/db"
	"github.com/Nyxoy/restAPI/models"
	"github.com/Nyxoy/restAPI/utils"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func HandlePassUpdate(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	var token = params["token"]
	var passBody models.UpdatePass

	if err0 := json.NewDecoder(r.Body).Decode(&passBody); err0 != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid body")
		log.Println("Invalid body")
		return
	}
	fmt.Println(token)
	fmt.Println(passBody)

	if token == "" {
		utils.WriteError(w, http.StatusBadRequest, "Missing token in the string")
		return
	}

	var upuser []models.Reset

	db := db.CreateRestyClient()
	resp, err := db.R().SetQueryParam("reset_token", "eq."+token).Get(viper.GetString("DB_BASE_URL") + "/rest/v1/password_reset_tokens")
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	fmt.Println(string(resp.Body()))
	fmt.Println(resp.StatusCode())
	if resp.StatusCode() == 200 {
		if err1 := json.Unmarshal(resp.Body(), &upuser); err1 != nil {
			utils.WriteError(w, http.StatusConflict, "Invalid data")
			return
		}
		actUser := upuser[0]
		fmt.Println(actUser)
		timestamp := actUser.Expire_Time
		formattedTimestamp := timestamp[:10] + "T" + timestamp[11:] + "Z"
		timeToExpire, er := time.Parse(time.RFC3339, formattedTimestamp)
		if er != nil {
			utils.WriteError(w, http.StatusInternalServerError, "An error occured while parsing the time")
			log.Printf("An error occured while parsing the time %s for user id %d", timestamp, actUser.User_ID)
			return
		}
		if time.Now().After(timeToExpire) {
			utils.WriteError(w, http.StatusBadRequest, "The token has expired")
			return
		}
		fmt.Println(actUser.Used)
		if (!actUser.Used) && (actUser.Reset_token == token) {
			fmt.Println("Token matched")
			hashedPass, err := utils.Encrypt(passBody.Password)
			if err != nil {
				log.Println("An error occured during hashing ")
				return
			}
			data := map[string]interface{}{
				"password": hashedPass,
			}
			_, err1 := db.R().SetBody(data).Patch(viper.GetString("DB_BASE_URL") + "/rest/v1/users?id=eq." + fmt.Sprintf("%d", actUser.User_ID))
			if err1 != nil {
				utils.WriteError(w, http.StatusInternalServerError, "Error occured updating the databse")
				log.Printf("An error occured when updating the password for %d user", actUser.User_ID)
				return
			} else {
				utils.WriteError(w, http.StatusOK, "Password Updated")
			}

			resp2, err2 := db.R().SetBody(map[string]interface{}{
				"used": true,
			}).Patch(viper.GetString("DB_BASE_URL") + "/rest/v1/password_reset_tokens?reset_token=eq." + token)

			if err2 != nil {
				utils.WriteError(w, http.StatusInternalServerError, "Error occured updating the used databse")
				log.Printf("An error occured when updating the password for boolean updation")
				return
			}

			fmt.Println(resp2.StatusCode())

		} else {
			utils.WriteError(w, http.StatusBadRequest, "The token is used !")
			return
		}

	}
}
