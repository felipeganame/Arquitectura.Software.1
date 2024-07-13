package controllers

import (
	"backend/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RegisterC(c *gin.Context) {
	var userDetails struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&userDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	log.Println("Registering user:", userDetails.Username)

	var tipo string = "normal"
	err := services.RegisterS(userDetails.Username, userDetails.Password, tipo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
