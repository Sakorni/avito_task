package main

import (
	"avito_task/server"
	"log"
)

func main() {
	handler := server.Handler{}

	server := handler.InitRoutes()
	log.Fatalf("%v", server.Run(":8080"))
}
