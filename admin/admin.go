package admin

import (
	"net/http"

	"github.com/Nyxoy/restAPI/models"
	"github.com/Nyxoy/restAPI/utils"
)

// The admin should be able to access the product tables
// I modified the role  in the middleware and i am attaching the role in the request using the claims

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	mp := r.Context().Value("claims").(*models.Claims)

	if mp.UserType != "ADMIN" {
		utils.WriteError(w, http.StatusUnauthorized, "Not an admin")
		return
	} else {
		utils.WriteError(w, http.StatusOK, "Admin ")
		return
	}

}
