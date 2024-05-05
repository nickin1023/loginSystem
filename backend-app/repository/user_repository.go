package repository

import (
	model "backend-app/dbModel"
	"database/sql"
)

type userRepository struct {
	*sql.DB
}

type IUserRepository interface {
	GetUsers() (*model.User, error)
	// CreateUser(user *db.User) (string, error)
	// Login(uID, password string) (string, error)
}

func NewUserRepository(Conn *sql.DB) IUserRepository {
	return &userRepository{Conn}
}

func (repo *userRepository) GetUsers() (*model.User, error) {
	dbUser := model.User{}
	// start := repo.ITimer.Start()
	// err := repo.DB.QueryRow("SELECT CityName FROM [Weather].[dbo].[Users] WHERE AccessToken = (?)", token).Scan(&dbUser.CityName)
	// if err != nil {
	// 	return nil, repo.IRepositoryError.SelectFailed(err)
	// }
	// queryCTX := "SELECT CityName FROM [Weather].[dbo].[Users]"
	// end := repo.ITimer.End()
	// repo.ITimer.Result(queryCTX, start, end)
	return &dbUser, nil
}
