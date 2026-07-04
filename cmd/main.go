package main

import (
	"log"
	"net/http"
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

	ticketRepo := repository.NewTicketRepository(config.DB)
	ticketService := services.NewTicketService(ticketRepo)
	ticketHandler := handlers.NewTicketHandler(ticketService)

	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	routes.RegisterRoutes(router, &routes.Handlers{
		AuthHandler: authHandler,
		TicketHandler: ticketHandler,
	})

	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}
