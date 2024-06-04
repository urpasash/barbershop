package controllers

import (
	"barbershop/config"
	"barbershop/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegistrationAppointment struct {
	Username string `json:"username"`
	Time     string `json:"time"`
	Master   string `json:"master"`
}

func CreateAppointment(c *gin.Context) {
	var regApp RegistrationAppointment

	if err := c.ShouldBindJSON(&regApp); err != nil {
		log.Println("Error in parsing JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid data format"})
		return
	}

	var app models.Appointment
	if err := config.DB.Where("time = ?", regApp.Time).First(&app).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"success": false, "message": "Запись уже есть"})
		return
	} else if err != gorm.ErrRecordNotFound {
		log.Println("Database error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Database error"})
		return
	}
	newApp := models.Appointment{
		Username: regApp.Username,
		Time:     regApp.Time,
		Master:   regApp.Master,
	}

	if err := config.DB.Create(&newApp).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, app)
}

func GetAppointment(c *gin.Context) {
	var appointments []models.Appointment
	if err := config.DB.Find(&appointments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, appointments)
}

func DeleteAppointment(c *gin.Context) {
	appointmentIDStr := c.Param("id")

	appointmentID, err := strconv.Atoi(appointmentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid  ID"})
		return
	}

	var appointment models.Appointment
	if err := config.DB.First(&appointment, appointmentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}

	if err := config.DB.Delete(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Appointment deleted"})
}

func UpdateAppointment(c *gin.Context) {
	appointmentsID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var appointment models.Appointment
	if err := config.DB.Where("ID = ?", appointmentsID).First(&appointment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": " not found"})
		return
	}
	fmt.Println(appointment)

	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	if err := config.DB.Save(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointment) //
}
