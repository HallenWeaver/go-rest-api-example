package service

import (
	"alexandre/gorest/app/model"
	"alexandre/gorest/app/repository"
	"context"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (us *UserService) CreateUser(ctx context.Context, newUser model.User, role model.UserRole) (*model.User, error) {
	return us.UserRepository.CreateUser(ctx, newUser, role)
}

func (us *UserService) LoginUser(ctx context.Context, loginUser model.TokenRequest) (*model.User, error) {
	return us.UserRepository.LoginUser(ctx, loginUser)
}
