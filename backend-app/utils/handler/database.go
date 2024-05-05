package handler

import (
	"database/sql"
	"log/slog"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"

	"backend-app/config"
)

var Db *sql.DB

func OpenDBConnection(config config.Config) (*sql.DB, error) {
	user := config.Db.User
	pw := config.Db.Pass
	db_name := config.Db.DatabaseName
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
		return nil, err
	}
	checkConnect(3)

	slog.Info("DB connected.")

	return Db, nil
}

func checkConnect(count int) error {
	var err error
	if err = Db.Ping(); err != nil {
		if count == 0 {
			return err
		}
		time.Sleep(time.Second * 1)
		count--
		slog.Info("message", "retry... count:%v", strconv.Itoa(count))
		if e := checkConnect(count); e != nil {
			return e
		}
	}

	return nil
}
