package user

import (
	"encoding/json"
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

}
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//Get The JSON payload
	var user models.User
	if err := utils.ParseJSON(r, &user, w); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
	}

	db := db.CreateRestyClient()
	ctx := r.Context()
	resp, err := db.R().
		SetContext(ctx).
		SetQueryParam("email", "eq."+user.Email).      // Set the query parameter for email
		Get(viper.GetString("DB_BASE_URL") + "/users") // Correctly access the BaseURL constant from db package

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Internal Server error")
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
	data := map[string]interface{}{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	}
	resp1, err1 := db.R().SetBody(data).Post(viper.GetString("DB_BASE_URL") + "/users")
	if err1 != nil {
		http.Error(w, "Failed to register the user", http.StatusInternalServerError)
		log.Println("Failed to register the user ")
		return
	}
	if resp1.StatusCode() >= 200 && resp1.StatusCode() < 300 {
		utils.WriteError(w, resp1.StatusCode(), "The user is registered successfully")
	}

}
