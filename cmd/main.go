package main

import (
	"log"
	"os"

	"github.com/Divyshekhar/eva-bharat-assignment/internal/config"
	"github.com/Divyshekhar/eva-bharat-assignment/internal/handlers"
	"github.com/Divyshekhar/eva-bharat-assignment/internal/repository"
	"github.com/Divyshekhar/eva-bharat-assignment/internal/routes"
	"github.com/Divyshekhar/eva-bharat-assignment/internal/services"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnv()
	config.ConnectDb()
	// config.MigrateDB()
}

func main() {

	userRepo := repository.NewUserRepository(config.DB)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	router := gin.Default()

	routes.RegisterRoutes(router, &routes.Handlers{AuthHandler: authHandler})

	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}
