package repository

import (
	"auth_service/internal/model"
	"errors"

	"gorm.io/gorm"
)

type AuthRepoMethod interface {
	CreateUser(user *model.User)error
	FindUser(username string)(*model.User, error)
}
type AuthRepo struct {
	DB *gorm.DB
}
func NewAuthRepo(db *gorm.DB)*AuthRepo{
	return &AuthRepo{DB:db}
}
func(ar *AuthRepo)CreateUser(user *model.User)error{
	var existingUser model.User
	if err := ar.DB.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return errors.New("username exists")
	}
	if err := ar.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return errors.New("email exists")
	}
	if err := ar.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
func(ar *AuthRepo)FindUser(username string)(*model.User, error){
	var existingUser model.User
	if err := ar.DB.Where("username = ?", username).First(&existingUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("username not found")
		} else {
			return nil, err
		}
	}
	return &existingUser, nil
}