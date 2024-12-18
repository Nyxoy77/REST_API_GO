package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Nyxoy/restAPI/db"
	"github.com/Nyxoy/restAPI/models"
	"github.com/Nyxoy/restAPI/utils"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/login", h.handleLogin).Methods("POST")
	r.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.Login
	if err := utils.ParseJSON(r, &user, w); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Error retrieving the data")
	}
	fmt.Println(user)
	db := db.CreateRestyClient()
	resp, err := db.R().SetQueryParam("email", "eq."+user.Email).Get(viper.GetString("DB_BASE_URL") + "/rest/v1/users")
	if err != nil {
		log.Printf("An error occured at the time of logging %s", err)
		http.Error(w, fmt.Sprintf("An error occurred at the time of logging: %s", err), http.StatusInternalServerError)
		return
	}
	var respUser []models.User
	err1 := json.Unmarshal(resp.Body(), &respUser)
	if err1 != nil {
		log.Println("An error occured at the time of parsing the data")
		http.Error(w, "An error occured at the time of parsing the data for logging", http.StatusInternalServerError)
		return
	}
	var actUser = respUser[0]
	fmt.Println(actUser.Password, " ", user.Password)
	matches := utils.CheckHashPass(user.Password, actUser.Password)
	fmt.Println(matches)

	if matches {
		utils.WriteError(w, http.StatusFound, "The user successfully logged in")
	} else {
		utils.WriteError(w, http.StatusFound, "The user successfully not logged in")
	}
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//Get The JSON payload
	var user models.User
	if err := utils.ParseJSON(r, &user, w); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
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
