package service

import (
	"auth_service/internal/model"
	"auth_service/internal/repository"
	"auth_service/internal/utils"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceMethod interface {
	RegisterService(username, email, password string) error
	Login(username, password string)(*model.User, string, error)
}
type AuthService struct {
	Repo	repository.AuthRepoMethod
}
func NewAuthService(repo repository.AuthRepoMethod)*AuthService{
	return &AuthService{Repo: repo}
}
func(as *AuthService)RegisterService(username, email, password string)error{
	if err := utils.CheckValidForm(username, email, password); err != nil {
		return err
	}
	hashedPassword, err := utils.GenerateHashPassword(password)
	if err != nil {
		return err
	}
	user := &model.User{
		UserID: utils.GenerateID(email),
		Username: username,
		Email: email,
		Password: hashedPassword,
		CreatedAt: time.Now(),
	}
	if err := as.Repo.CreateUser(user); err != nil {
		return errors.New("failed to create user: " + err.Error())
	}
	return nil
}
func(as *AuthService)Login(username, password string)(*model.User, string, error){
	user, err := as.Repo.FindUser(username)
	if err != nil {
		return nil, "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil{
		return nil, "", errors.New("invalid password")
	}
	token, err := generateJWT(username)
	if err != nil {
		return nil, "", errors.New("failed to create token")
	}
	return user, token, nil
}


func generateJWT(username string)(string, error){
	var jwtKey = os.Getenv("JWT_KEY")
	type claims struct {
		Username string 
		jwt.RegisteredClaims
	}

	exp := time.Now().Add(24 * time.Hour)

	claim := &claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}