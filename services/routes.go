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

}

func (h *Handler) RegisterProtectedRoutes(r *mux.Router) {
	r.HandleFunc("/products", VerifyJWT(FetchProducts)).Methods("GET")
}
