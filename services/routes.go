package services

import (
	"github.com/Nyxoy/restAPI/admin"
	"github.com/Nyxoy/restAPI/user_customer"
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

func (h *Handler) RegisterAdminRoutes(r *mux.Router) {
	r.HandleFunc("/get_all_users", VerifyJWT(admin.AdminHandler(admin.GetAllUsers))).Methods("GET")
	r.HandleFunc("/get_admins", VerifyJWT(admin.AdminHandler(admin.GetAllAdmins))).Methods("GET")
	r.HandleFunc("/get_customers", VerifyJWT(admin.AdminHandler(admin.GetallCustomers))).Methods("GET")
	r.HandleFunc("/add_product", VerifyJWT(admin.AdminHandler(admin.AddProduct))).Methods("POST")
	r.HandleFunc("/remove_product", VerifyJWT(admin.AdminHandler(admin.RemoveProduct))).Methods("DELETE")
	r.HandleFunc("/update_price", VerifyJWT(admin.AdminHandler(admin.UpdatePrice))).Methods("PUT")
}

func (h *Handler) RegisterUserRoutes(r *mux.Router) {
	r.HandleFunc("/add_to_cart", VerifyJWT(admin.UserHandler(user_customer.AddItem))).Methods("POST")
	r.HandleFunc("/remove_item", VerifyJWT(admin.UserHandler(user_customer.RemoveItem))).Methods("DELETE")

}
