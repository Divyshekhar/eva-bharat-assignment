package handlers

import (
	"net/http"

	"github.com/Divyshekhar/eva-bharat-assignment/internal/dto"
	"github.com/Divyshekhar/eva-bharat-assignment/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct{
	authService services.AuthService
}

func NewAuthHandler(authservice services.AuthService) *AuthHandler{
	return &AuthHandler{authService: authservice}
}

func (h *AuthHandler) Register(c *gin.Context){
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := h.authService.Register(req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func (h *AuthHandler) Login(c *gin.Context){
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	token, err := h.authService.Login(req); 
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, token)
}