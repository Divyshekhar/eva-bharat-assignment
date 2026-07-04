package main

import (
	"net/http"
	"os"

	"github.com/Divyshekhar/eva-bharat-assignment/internal/config"
	"github.com/gin-gonic/gin"
)

func init(){
	config.LoadEnv()
	config.ConnectDb()
	// config.MigrateDB()
}

func main(){
	router := gin.Default()
	router.GET("/health", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is healthy",
		})
	})
	router.Run(":"+os.Getenv("PORT"))
}