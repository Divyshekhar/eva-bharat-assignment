package routes

import (
	"net/http"

	"github.com/Divyshekhar/eva-bharat-assignment/internal/handlers"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	AuthHandler *handlers.AuthHandler
	TicketHandler *handlers.TicketHandler
}

func RegisterRoutes(router *gin.Engine, h *Handlers){
	router.GET("/health", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	RegisterAuthRoutes(router, h.AuthHandler)
	RegisterTicketRoutes(router, h.TicketHandler)
}
