package service

import (
	"alexandre/gorest/app/model"
	"alexandre/gorest/app/repository"
	"context"
	"testing"
)

// Defining Mock User Repositories

type MockUserRepository struct{}

type MockUserRepositoryWithErrors struct{}

func (mur *MockUserRepository) CreateUser(ctx context.Context, newUser model.User, role model.UserRole) (*model.User, error) {
	return &newUser, nil
}

func (mur *MockUserRepository) LoginUser(ctx context.Context, loginUser model.TokenRequest) (*model.User, error) {
	return &model.User{
		Username: loginUser.Username,
		Password: loginUser.Password,
	}, nil
}

func (murwe *MockUserRepositoryWithErrors) CreateUser(ctx context.Context, newUser model.User, role model.UserRole) (*model.User, error) {
	return nil, ErrorMessageTesting
}

func (murwe *MockUserRepositoryWithErrors) LoginUser(ctx context.Context, loginUser model.TokenRequest) (*model.User, error) {
	return nil, ErrorMessageTesting
}

type MockUserRepositoryWrapper struct {
	userRepo       repository.IUserRepository
	isValidTesting bool
}

// Building Actual Tests
func TestUserService(t *testing.T) {
	userRepos := []MockUserRepositoryWrapper{
		{userRepo: &MockUserRepository{}, isValidTesting: true},
		{userRepo: &MockUserRepositoryWithErrors{}, isValidTesting: false},
	}

	for _, userRepo := range userRepos {
		userService := NewUserService(userRepo.userRepo)

		createdUser, err := userService.CreateUser(context.Background(), model.User{}, model.Standard)
		AssertValidityCondition(t, createdUser, err, userRepo.isValidTesting)

		loginUser, err := userService.LoginUser(context.Background(), model.TokenRequest{})
		AssertValidityCondition(t, loginUser, err, userRepo.isValidTesting)

	}
}
