package handler

import (
	"alexandre/gorest/app/helper"
	"alexandre/gorest/app/model"
	"alexandre/gorest/app/service"
	"context"
)

type MockUserService struct{}

type MockUserServiceWithErrors struct{}

func (mus *MockUserService) CreateUser(ctx context.Context, newUser model.User, role model.UserRole) (*model.User, error) {
	return &model.User{}, nil
}

func (mus *MockUserService) LoginUser(ctx context.Context, loginUser model.TokenRequest) (*model.User, error) {
	return &model.User{}, nil
}

func (muswe *MockUserServiceWithErrors) CreateUser(ctx context.Context, newUser model.User, role model.UserRole) (*model.User, error) {
	return nil, helper.ErrorMessageTesting
}

func (muswe *MockUserServiceWithErrors) LoginUser(ctx context.Context, loginUser model.TokenRequest) (*model.User, error) {
	return nil, helper.ErrorMessageTesting
}

type MockUserTestCaseWrapper struct {
	userService        service.IUserService
	isValidTesting     bool
	userPayload        interface{}
	ExpectedStatusCode int
}
