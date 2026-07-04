package routes

import (
	"github.com/Divyshekhar/eva-bharat-assignment/internal/handlers"
	"github.com/Divyshekhar/eva-bharat-assignment/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterTicketRoutes(router *gin.Engine, ticketHandler *handlers.TicketHandler) {
	ticket := router.Group("/tickets")
	ticket.Use(middleware.AuthMiddleware())
	{
		ticket.POST("/", ticketHandler.Create)
		ticket.GET("/", ticketHandler.GetAll)
		ticket.GET("/:id", ticketHandler.GetByID)
		ticket.PATCH("/:id/status", ticketHandler.UpdateStatus)
	}
}
