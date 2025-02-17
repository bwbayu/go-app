package main

import (
	"fmt"
	"go-app/internal/config"
	"go-app/internal/handler"
	"go-app/internal/repository"
	"go-app/internal/routes"
	"go-app/internal/service"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error load .env file")
	}
	config.ConnectDB()
	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	r := routes.SetupRouter(userHandler)

	fmt.Printf("Server running on port %v", os.Getenv("SERVER_PORT"))
	port := os.Getenv("SERVER_PORT")
	r.Run(fmt.Sprintf(":%s", port))
}
