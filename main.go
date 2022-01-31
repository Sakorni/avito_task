package main

import (
	"avito_task/repository"
	"avito_task/server"
	"avito_task/service"
	"flag"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	loadEnv()
	dbCfg := repository.DBConfig{
		DBName:   os.Getenv("DB_NAME"),
		Address:  os.Getenv("DB_ADDRESS"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	db, err := repository.InitDB(dbCfg)
	if err != nil {
		logrus.Fatalf("error occured during connecting to database: %v", err)
	}
	rep := repository.NewBalanceRepository(db)
	service := service.NewBalanceService(rep)
	handler := server.NewHandler(service)
	engine := handler.InitRoutes()
	logrus.Fatalf("%v", engine.Run(os.Getenv("AVITO_PORT")))
}

func loadEnv() {
	var envDirectory string
	flag.StringVar(&envDirectory, "env_dir", "./config",
		"Path to directory, which contains configuration files")
	flag.Parse()
	dirContent, err := ioutil.ReadDir(envDirectory)
	if err != nil {
		logrus.Fatal(err)
	}
	files := make([]string, 0)
	absPath, _ := filepath.Abs(envDirectory)
	for _, file := range dirContent {
		if !file.IsDir() {
			files = append(files, filepath.Join(absPath, file.Name()))
		}
	}

	err = godotenv.Load(files...)
	if err != nil {
		logrus.Fatalf("can't load env files: %v", err)
	}
}
