package service

import (
	"auth_service/internal/model"
	"errors"
	"testing"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Mock repository
type MockAuthRepo struct {
	users map[string]*model.User
}

func (m *MockAuthRepo) CreateUser(user *model.User) error {
	m.users[user.Username] = user
	return nil
}

func (m *MockAuthRepo) FindUser(username string) (*model.User, error) {
	user, exists := m.users[username]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// Test RegisterService
func TestRegisterService(t *testing.T) {
	mockRepo := &MockAuthRepo{users: make(map[string]*model.User)}
	authService := NewAuthService(mockRepo)

	// Test successful registration
	err := authService.RegisterService("testuser", "test@example.com", "Password_123")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify user was created
	user, err := mockRepo.FindUser("testuser")
	if err != nil || user.Username != "testuser" {
		t.Fatalf("expected user to be created, got %v", err)
	}

	// Test registration with invalid email
	err = authService.RegisterService("testuser2", "invalid-email", "password123")
	if err == nil {
		t.Fatal("expected an error for invalid email, got none")
	}
}

// Test Login
func TestLogin(t *testing.T) {
	mockRepo := &MockAuthRepo{users: make(map[string]*model.User)}
	authService := NewAuthService(mockRepo)

	// Setup a user for testing
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := &model.User{
		UserID:   "1",
		Username: "testuser",
		Email:    "test@example.com",
		Password: string(hashedPassword),
		CreatedAt: time.Now(),
	}
	mockRepo.users[user.Username] = user

	// Test successful login
	loggedInUser, token, err := authService.Login("testuser", "password123")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if loggedInUser.Username != "testuser" {
		t.Fatalf("expected username 'testuser', got '%s'", loggedInUser.Username)
	}
	if token == "" {
		t.Fatal("expected a token, got none")
	}

	// Test login with incorrect password
	_, _, err = authService.Login("testuser", "wrongpassword")
	if err == nil {
		t.Fatal("expected an error for incorrect password, got none")
	}

	// Test login for a non-existing user
	_, _, err = authService.Login("nonexistent", "password123")
	if err == nil {
		t.Fatal("expected an error for non-existent user, got none")
	}
}
