package main

import (
	"avito_task/repository"
	"avito_task/server"
	"avito_task/service"
	"github.com/sirupsen/logrus"
)

func main() {
	dbCfg := repository.DBConfig{
		DBName:   "avito",
		Address:  "localhost:3306",
		Username: "root",
		Password: "admin",
	}
	db, err := repository.InitDB(dbCfg)
	if err != nil {
		logrus.Fatalf("error occured while opening database: %v", err)
	}
	rep := repository.NewBalanceRepository(db)
	service := service.NewBalanceService(rep)
	handler := server.NewHandler(service)

	engine := handler.InitRoutes()
	logrus.Fatalf("%v", engine.Run(":8080"))
}
