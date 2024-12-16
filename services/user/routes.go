package user

import (
	"encoding/json"
	"fmt"
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

}
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//Get The JSON payload
	var user models.User
	if err := utils.ParseJSON(r, &user, w); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// fmt.Println(user)
	// Check if the user already exists in the database
	db := db.CreateRestyClient()
	resp, err := db.R().
		SetQueryParam("email", "eq."+user.Email).      // Set the query parameter for email
		Get(viper.GetString("DB_BASE_URL") + "/users") // Correctly access the BaseURL constant from db package

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	fmt.Println(string(resp.Body()))
	// fmt.Println(resp.StatusCode())
	if resp.StatusCode() == 200 {
		var users []models.User
		if err := json.Unmarshal(resp.Body(), &users); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		if len(users) > 0 {

			fmt.Println("The user is already registered. Please Login")
			return
		}
	}

}
