package repository

import (
	"auth_service/internal/model"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)
type MockRepoMethod struct{
	mock.Mock
}
func(m *MockRepoMethod)CreateUser(user *model.User)error{
	args := m.Called(user)
	return args.Error(0)
}
func TestCreateUser(t *testing.T) {
	mockRepo := new(MockRepoMethod)
	user := &model.User{
		Username: "test",
		Email: "test@gmail.com",
		Password: "test1234",
	}
	createUserMock := mockRepo.On("CreateUser", user)
	createUserMock.Return(nil).Once()
	err := mockRepo.CreateUser(user)
	assert.Nil(t, err)

	createUserMock.Return(errors.New("username exists")).Once()
	err = mockRepo.CreateUser(user)
	assert.NotNil(t, err)
	assert.Equal(t, "username exists", err.Error())

	createUserMock.Return(errors.New("email exists")).Once()
	err = mockRepo.CreateUser(user)
	assert.NotNil(t, err)
	assert.Equal(t, "email exists", err.Error())
}