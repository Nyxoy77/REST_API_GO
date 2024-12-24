package admin

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Nyxoy/restAPI/db"
	"github.com/Nyxoy/restAPI/models"
	"github.com/Nyxoy/restAPI/utils"
	"github.com/spf13/viper"
)

// The admin should be able to access the product tables
// I modified the role  in the middleware and i am attaching the role in the request using the claims

func AdminHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mp := r.Context().Value("claims").(*models.Claims)
		if mp.UserType != "ADMIN" {
			utils.WriteError(w, http.StatusUnauthorized, "Not an admin")
			return
		}
		next(w, r)
	}
}
func UserHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mp := r.Context().Value("claims").(*models.Claims)
		if mp.UserType != "USER" {
			utils.WriteError(w, http.StatusUnauthorized, "Not a user")
			return
		}
		next(w, r)
	}
}

// This function will only admins and will give complete control over the users

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := db.CreateRestyClient().R().SetQueryParam("select", "*").Get(viper.GetString("DB_BASE_URL") + "/rest/v1/users")
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error retrieving the users data")
		log.Println("An error retrieving the users data ")
		return
	}
	var users []models.User
	if resp.StatusCode() == 200 {
		if err := json.Unmarshal(resp.Body(), &users); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, "Error parsing the data ")
			log.Println("An error parsing the users data ")
			return
		}
		if len(users) == 0 {
			utils.WriteError(w, http.StatusOK, "No users exist")
			return
		} else {
			for _, value := range users {
				encoder := json.NewEncoder(w)
				encoder.SetIndent("", "\n") // Use 4 spaces for indentation (or adjust as needed)
				encoder.Encode(value)
			}
		}
	}
}

func GetAllAdmins(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := db.CreateRestyClient().R().SetQueryParam("user_type", "eq."+"ADMIN").Get(viper.GetString("DB_BASE_URL") + "/rest/v1/users")
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error retrieving the users data")
		log.Println("An error retrieving the users data ")
		return
	}
	var users []models.User
	if resp.StatusCode() == 200 {
		if err := json.Unmarshal(resp.Body(), &users); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, "Error parsing the data ")
			log.Println("An error parsing the users data ")
			return
		}
		if len(users) == 0 {
			utils.WriteError(w, http.StatusOK, "No users exist")
			return
		} else {
			for _, value := range users {
				encoder := json.NewEncoder(w)
				encoder.SetIndent("", "\n") // Use 4 spaces for indentation (or adjust as needed)
				encoder.Encode(value)
			}
		}
	}
}

func GetallCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := db.CreateRestyClient().R().SetQueryParam("user_type", "eq."+"USER").Get(viper.GetString("DB_BASE_URL") + "/rest/v1/users")
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error retrieving the users data")
		log.Println("An error retrieving the users data ")
		return
	}
	var users []models.User
	if resp.StatusCode() == 200 {
		if err := json.Unmarshal(resp.Body(), &users); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, "Error parsing the data ")
			log.Println("An error parsing the users data ")
			return
		}
		if len(users) == 0 {
			utils.WriteError(w, http.StatusOK, "No users exist")
			return
		} else {
			for _, value := range users {
				encoder := json.NewEncoder(w)
				encoder.SetIndent("", "\n") // Use 4 spaces for indentation (or adjust as needed)
				encoder.Encode(value)
			}
		}
	}
}
