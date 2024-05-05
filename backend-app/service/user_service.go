package service

import (
	"backend-app/model"
	"backend-app/repository"
	"log/slog"
)

type userService struct {
	repository.IUserRepository
}

type IUserService interface {
	GetUsersService() ([]model.User, error)
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &userService{repo}
}

func (repo *userService) GetUsersService() ([]model.User, error) {
	users := []model.User{}
	dbUsers, err := repo.GetUsers()
	if err != nil {
		return nil, err
	}
	for _, dbUser := range dbUsers {
		user := model.User{}
		user.Id = dbUser.Id
		user.Name = dbUser.Name
		user.Role = dbUser.Role
		user.Email = dbUser.Email
		user.CreatedDate = dbUser.CreatedDate
		user.UpdatedDate = dbUser.UpdatedDate
		users = append(users, user)
	}
	slog.Info("message:", "users", users)
	return users, nil
}
