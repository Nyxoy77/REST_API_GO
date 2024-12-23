package services

import (
	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/login", HandleLogin).Methods("POST")
	r.HandleFunc("/register", HandleRegister).Methods("POST")
	r.HandleFunc("/forgot", HandleForgot).Methods("POST")
	r.HandleFunc("/update_pass/{token}", HandlePassUpdate)
	r.HandleFunc("/refresh", RefreshTokenHandler)

}

func (h *Handler) RegisterProtectedRoutes(r *mux.Router) {
	r.HandleFunc("/products", VerifyJWT(FetchProducts)).Methods("GET")
}
