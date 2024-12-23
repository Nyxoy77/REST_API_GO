package api

import (
	"log"
	"net/http"

	"github.com/Nyxoy/restAPI/services"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	userHandler := services.NewHandler()
	userHandler.RegisterRoutes(subrouter)
	protectedSubRouter := router.PathPrefix("/api/v1/protected").Subrouter()
	adminSubRouter := router.PathPrefix("/api/v1/protected/admin").Subrouter()
	userHandler.RegisterProtectedRoutes(protectedSubRouter)
	userHandler.RegisterAdminRoutes(adminSubRouter)
	log.Println("Listening on ", s.addr)
	return http.ListenAndServe(s.addr, router)
}
