package routes

import (
	"github.com/Divyshekhar/eva-bharat-assignment/internal/handlers"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	AuthHandler   *handlers.AuthHandler
	TicketHandler *handlers.TicketHandler
}

func RegisterRoutes(router *gin.Engine, h *Handlers) {

	RegisterAuthRoutes(router, h.AuthHandler)
	RegisterTicketRoutes(router, h.TicketHandler)
}
