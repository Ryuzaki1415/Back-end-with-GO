package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {   
}

func NewHandler() *Handler {     // we return the address of the handler created
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {  //we take in the subrouter from  api.go 


	//services provided

	router.HandleFunc("/login",h.handleLogin).Methods("POST")  // here we can have all different functionalities.
	router.HandleFunc("/register",h.handleRegister).Methods("POST")
}



func (h *Handler)handleLogin(w http.ResponseWriter,r *http.Request){  // logic for handling login

}
func (h *Handler)handleRegister(w http.ResponseWriter,r *http.Request){  //login for handling Registration 

}