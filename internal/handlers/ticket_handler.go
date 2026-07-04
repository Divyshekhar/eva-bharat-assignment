package handlers

import (
	"net/http"

	"github.com/Divyshekhar/eva-bharat-assignment/internal/dto"
	"github.com/Divyshekhar/eva-bharat-assignment/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TicketHandler struct {
	service services.TicketService
}

func NewTicketHandler(service services.TicketService) *TicketHandler {
	return &TicketHandler{service: service}
}

func (h *TicketHandler) Create(c *gin.Context) {
	var req dto.CreateTicketRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	userId := c.MustGet("userID").(uuid.UUID)
	err := h.service.Create(userId, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Ticket created successfully",
	})
}

func (h *TicketHandler) GetAll(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)
	tickets, err := h.service.GetAll(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tickets)
}

func (h *TicketHandler) GetByID(c *gin.Context) {
	userID := c.MustGet("userID").(uuid.UUID)
	ticketID, err := uuid.Parse(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ticket id",
		})
		return
	}
	ticket, err := h.service.GetById(userID, ticketID)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ticket)
}

func (h *TicketHandler) UpdateStatus(c *gin.Context){
	userID := c.MustGet("userID").(uuid.UUID)
	ticketID, err := uuid.Parse(c.Param("id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ticket id",
		})
		return
	}
	var req dto.UpdateTicketRequest
	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	er := h.service.UpdateStatus(userID, ticketID, req.Status)
	if er != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": er.Error(),
		})
		return 
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Ticket updates successfully",
	})
}
