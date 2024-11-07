package api

import (
	"auth_service/internal/service"
	"encoding/json"
	"net/http"
)

type AuthHandlerMethod interface{
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
}
type AuthHandler struct{
	Service service.AuthServiceMethod
}
func NewAuthHandler(service *service.AuthService)*AuthHandler{
	return &AuthHandler{Service: service}
}
func(ah *AuthHandler)Login(w http.ResponseWriter, r *http.Request){
	var loginData struct {
		Username	string	`json:"username"`
		Password	string	`json:"password"`
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, "Bad request: " + err.Error(), http.StatusInternalServerError)
		return
	}
	user, token, err := ah.Service.Login(loginData.Username, loginData.Password)
	if err != nil {
		http.Error(w, "Failed: " + err.Error(), http.StatusUnauthorized)
		return
	}
	res := map[string]interface{}{
		"message": "login successful",
		"user": user.Username,
		"token": token,
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Bad request: " + err.Error(), http.StatusInternalServerError)
		return
	}
}

func(ah *AuthHandler)Register(w http.ResponseWriter, r *http.Request){
	var user struct {
		Username	string	`json:"username"`
		Email		string	`json:"email"`
		Password	string	`json:"password"`
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad request: " + err.Error(), http.StatusBadRequest)
		return
	}
	if err := ah.Service.RegisterService(user.Username, user.Email, user.Password); err != nil {
		http.Error(w, "Bad request: " + err.Error(), http.StatusInternalServerError)
		return
	}
	res := map[string]string{
		"message": "register successful",
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}