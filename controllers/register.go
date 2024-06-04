package controllers

import (
	"barbershop/config"
	"barbershop/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegistrationData struct {
	UsernameReg string `json:"username_reg" binding:"required"`
	PasswordReg string `json:"password_reg" binding:"required"`
}

func Register(c *gin.Context) {
	log.Println("Received registration request")

	var registrationData RegistrationData

	if err := c.ShouldBindJSON(&registrationData); err != nil {
		log.Println("Error in parsing JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid data format"})
		return
	}

	var existingUser models.User
	if err := config.DB.Where("username = ?", registrationData.UsernameReg).First(&existingUser).Error; err == nil {
		log.Println("Username already exists:", registrationData.UsernameReg) // Логируем, если имя занято
		c.JSON(http.StatusConflict, gin.H{"success": false, "message": "Username already taken"})
		return
	} else if err != gorm.ErrRecordNotFound {
		log.Println("Database error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Database error"})
		return
	}

	newUser := models.User{
		Username: registrationData.UsernameReg,
		Password: registrationData.PasswordReg,
		IsAdmin:  false,
	}

	if err := config.DB.Create(&newUser).Error; err != nil {
		log.Println("Error creating user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error creating user"})
		return
	}

	log.Println("Registration successful:", newUser.Username)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Registration successful", "user": newUser})
}
