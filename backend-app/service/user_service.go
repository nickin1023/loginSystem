package service

import (
	"backend-app/model"
	"backend-app/repository"
)

type userService struct {
	repository.IUserRepository
}

type IUserService interface {
	GetUsers() (*model.User, error)
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &userService{repo}
}

func (repo *userService) GetUsers() (*model.User, error) {
	user := model.User{}
	return &user, nil
}
