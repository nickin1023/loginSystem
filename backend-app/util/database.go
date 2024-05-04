package util

import (
	"backend-app/config"
	"database/sql"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

type User struct {
	Id          int
	Name        string
	Password    string
	Role        string
	Email       string
	CreatedDate string
	UpdatedDate string
}

func init() {
	user := config.Config.Db.User
	pw := config.Config.Db.Pass
	db_name := config.Config.Db.DatabaseName
	dbConfig := mysql.Config{
		User:                 user,
		Passwd:               pw,
		Net:                  "tcp",
		Addr:                 "mysql",
		DBName:               db_name,
		ParseTime:            true,
		AllowNativePasswords: true,
		Collation:            "utf8mb4_bin",
	}
	var err error
	if Db, err = sql.Open("mysql", dbConfig.FormatDSN()); err != nil {
		slog.Error("Db open error:", err)
		os.Exit(1)
	}
	checkConnect(10)

	slog.Info("DB connected.")
}

func checkConnect(count int) {
	var err error
	if err = Db.Ping(); err != nil {
		slog.Warn("DB connection warning:", err)
		time.Sleep(time.Second * 1)
		count--
		slog.Info("message", "retry... count:%v", strconv.Itoa(count))
		checkConnect(count)
	}
}

func FindAll() {
	selected, err := Db.Query("SELECT * FROM USERS")
	if err != nil {
		slog.Error("Db find all error:", err)
		os.Exit(1)
	}
	for selected.Next() {
		user := User{}
		err = selected.Scan(&user.Id, &user.Name, &user.Password, &user.Role, &user.Email, &user.CreatedDate, &user.UpdatedDate)
		if err != nil {
			slog.Error("DB ERROR cannot get user", err)
			panic(err)
		}
		slog.Info("user", slog.Any("data", user))
	}
	selected.Close()
}
