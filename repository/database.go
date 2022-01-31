package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type DBConfig struct {
	Username string
	Password string
	Address  string
	DBName   string
}

func InitDB(config DBConfig) (*sql.DB, error) {
	cfgString := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Username, config.Password, config.Address, config.DBName)
	db, err := sql.Open("mysql", cfgString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	logrus.Info("Connected to database")
	return db, nil
}
