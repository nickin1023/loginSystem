package repository

import (
	model "backend-app/dbModel"
	"database/sql"
)

type userRepository struct {
	*sql.DB
}

type IUserRepository interface {
	GetUsers() ([]model.User, error)
}

func NewUserRepository(Conn *sql.DB) IUserRepository {
	return &userRepository{Conn}
}

func (repo *userRepository) GetUsers() ([]model.User, error) {
	selected, err := repo.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	users := []model.User{}
	for selected.Next() {
		user := model.User{}
		err = selected.Scan(&user.Id, &user.Name, &user.Password, &user.Role, &user.Email, &user.CreatedDate, &user.UpdatedDate)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	selected.Close()
	return users, nil
}
