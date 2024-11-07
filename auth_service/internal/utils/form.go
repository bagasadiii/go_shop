package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
	"unicode"

	form "github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

func CheckValidForm(username, email, password string) error {
	if len(username) < 3 {
		return errors.New("username must be at least 3 characters long")
	}
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	if !form.IsEmail(email){
		return errors.New("invalid email address")
	}
	if form.HasUpperCase(username){
		return errors.New("username cannot contain uppercase")
	}
	if form.HasUpperCase(email){
		return errors.New("email cannot contain uppercase")
	}
	if form.HasWhitespace(username) || form.HasWhitespaceOnly(username){
		return errors.New("username cannot contain space")
	}
	if form.HasWhitespace(email) || form.HasWhitespaceOnly(email){
		return errors.New("username cannot contain space")
	}
	if !strongPassword(password){
		return errors.New("password must contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	}
	for _, char := range username {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '_'{
			return errors.New("username only can contain letter, digit and underscore (_)")
		}
	}
	return nil
}
func strongPassword(password string)bool{
	var hasUpper, hasLower, hasNumber, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSpecial
}
func GenerateID(email string) string {
	timestamp := time.Now().UnixNano()
	input := fmt.Sprintf("%s%d", email, timestamp)
	hash := sha256.Sum256([]byte(input))
	id := hex.EncodeToString(hash[:])[:8]
	return id
}
func GenerateHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash password")
	}
	return string(hashedPassword), nil
}