package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Nyxoy/restAPI/db"
	"github.com/Nyxoy/restAPI/models"
	"github.com/Nyxoy/restAPI/utils"
	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	//Get The JSON payload
	var user models.User
	if err := utils.ParseJSON(r, &user, w); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
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
	db := db.CreateRestyClient()
	ctx := r.Context()
	resp, err := db.R().
		SetContext(ctx).
		SetQueryParam("email", "eq."+user.Email).              // Set the query parameter for email
		Get(viper.GetString("DB_BASE_URL") + "/rest/v1/users") // Correctly access the BaseURL constant from db package

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Internal Server error")
		return
	}

	if resp.StatusCode() == 200 {
		var users []models.User
		if err := json.Unmarshal(resp.Body(), &users); err != nil {
			utils.WriteError(w, http.StatusBadRequest, "Wrong data fetched")
			return
		}

		if len(users) > 0 {
			utils.WriteError(w, http.StatusConflict, "The user already exists!")
			return
		}
	}
	hashedPass, _ := utils.Encrypt(user.Password)
	data := map[string]interface{}{
		"firstname": user.FirstName,
		"lastname":  user.LastName,
		"email":     user.Email,
		"password":  hashedPass,
		"phone":     user.Phone,
		"user_type": user.UserType,
	}
	resp1, err1 := db.R().SetBody(data).Post(viper.GetString("DB_BASE_URL") + "/rest/v1/users")
	if err1 != nil {
		http.Error(w, "Failed to register the user", http.StatusInternalServerError)
		log.Println("Failed to register the user ")
		return
	}
	if resp1.StatusCode() >= 200 && resp1.StatusCode() < 300 {
		utils.WriteError(w, resp1.StatusCode(), "The user is registered successfully")
	}
	if resp1.StatusCode() >= 400 {
		log.Printf("Error: %s\n", resp.String())
		utils.WriteError(w, resp.StatusCode(), "Failed to store user data")
		log.Printf("Supabase Response Status: %d\n", resp.StatusCode())
		log.Printf("Supabase Response Body: %s\n", resp.String())
		return
	}

}
