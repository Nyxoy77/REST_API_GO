package logoutservice

import (
	"log"
	"net/http"
	"time"

	"github.com/Nyxoy/restAPI/caching"
	"github.com/Nyxoy/restAPI/utils"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	tokenstring := r.Header.Get("Authorization")
	if tokenstring == "" {
		utils.WriteError(w, http.StatusBadRequest, "Token empty")
		return
	}
	tokenstring = tokenstring[len("Bearer "):]

	if err := caching.BlackListToken(tokenstring, 24*time.Hour); err != nil {
		log.Printf("An error occured during token blacklisting %v ", err)
		return
	}
	utils.WriteError(w, http.StatusOK, "Logged out successfully")

}
