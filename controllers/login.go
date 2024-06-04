package controllers

import (
	"barbershop/config"
	"barbershop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var loginData LoginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid data format"})
		return
	}

	var user models.User
	if err := config.DB.Where("username = ? AND password = ?", loginData.Username, loginData.Password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid username or password"})
		return
	}
	
	if user.IsAdmin {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "Admin login successful", "is_admin": true, "user": user})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": true, "message": "User login successful", "is_admin": false, "user": user})
	}
}
