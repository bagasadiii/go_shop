package main

import (
	"auth_service/cmd/api"
	"auth_service/internal/config"
	"auth_service/internal/repository"
	"auth_service/internal/service"
	"log"
	"net/http"
)

func main() {
	DB := config.InitDBTest()
	authRepo := repository.NewAuthRepo(DB)
	authService := service.NewAuthService(authRepo)
	api := api.NewAuthHandler(authService)

	http.HandleFunc("/login", api.Login)
	http.HandleFunc("/register", api.Register)

	if err := http.ListenAndServe(":8000", nil); err != nil{
		log.Fatal("failed: ", err)
	}
}