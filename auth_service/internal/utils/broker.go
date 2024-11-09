package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendToUserService(id, email, username string)error{
	url := "http://localhost:8001/register"

	newUser := struct{
		ID       string `json:"userID"`
		Email    string `json:"email"`
		Username string `json:"username"`
	}{
		ID: id,
		Email: email,
		Username: username,
	}

	userJson, err := json.Marshal(newUser)
	if err != nil {
		return err
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(userJson))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send user data: %v", err)
	}
	return nil
}