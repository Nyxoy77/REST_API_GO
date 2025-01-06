package services

import (
	"github.com/Nyxoy/restAPI/admin"
	authservices "github.com/Nyxoy/restAPI/services/auth_services"
	forgotservices "github.com/Nyxoy/restAPI/services/forgot_services"
	myjwt "github.com/Nyxoy/restAPI/services/jwt_logic"
	logoutservice "github.com/Nyxoy/restAPI/services/logout_service"
	productservices "github.com/Nyxoy/restAPI/services/products"
	"github.com/Nyxoy/restAPI/user_customer"
	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/login", authservices.HandleLogin).Methods("POST")
	r.HandleFunc("/register", authservices.HandleRegister).Methods("POST")
	r.HandleFunc("/forgot", forgotservices.HandleForgot).Methods("POST")
	r.HandleFunc("/update_pass/{token}", forgotservices.HandlePassUpdate)
	r.HandleFunc("/refresh", myjwt.RefreshTokenHandler)
}

func (h *Handler) RegisterProtectedRoutes(r *mux.Router) {
	r.HandleFunc("/logout", myjwt.VerifyJWT(logoutservice.LogoutHandler)).Methods("POST")
	r.HandleFunc("/products", myjwt.VerifyJWT(productservices.FetchProducts)).Methods("GET")
}

func (h *Handler) RegisterAdminRoutes(r *mux.Router) {
	r.HandleFunc("/get_all_users", myjwt.VerifyJWT(admin.AdminHandler(admin.GetAllUsers))).Methods("GET")
	r.HandleFunc("/get_admins", myjwt.VerifyJWT(admin.AdminHandler(admin.GetAllAdmins))).Methods("GET")
	r.HandleFunc("/get_customers", myjwt.VerifyJWT(admin.AdminHandler(admin.GetallCustomers))).Methods("GET")
	r.HandleFunc("/add_product", myjwt.VerifyJWT(admin.AdminHandler(admin.AddProduct))).Methods("POST")
	r.HandleFunc("/remove_product", myjwt.VerifyJWT(admin.AdminHandler(admin.RemoveProduct))).Methods("DELETE")
	r.HandleFunc("/update_price", myjwt.VerifyJWT(admin.AdminHandler(admin.UpdatePrice))).Methods("PUT")
}

func (h *Handler) RegisterUserRoutes(r *mux.Router) {
	r.HandleFunc("/add_to_cart", myjwt.VerifyJWT(admin.UserHandler(user_customer.AddItem))).Methods("POST")
	r.HandleFunc("/remove_item", myjwt.VerifyJWT(admin.UserHandler(user_customer.RemoveItem))).Methods("DELETE")
	r.HandleFunc("/order", myjwt.VerifyJWT(admin.UserHandler(user_customer.AddToOrders))).Methods("POST")

}
