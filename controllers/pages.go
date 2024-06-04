package controllers

import (
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.HTML(200, "brb.html", gin.H{"Title": "Главная страница"})
}

func ServicesPage(c *gin.Context) {
	c.HTML(200, "services.html", gin.H{"Title": "Наши Услуги"})
}

func MastersPage(c *gin.Context) {
	c.HTML(200, "masters.html", gin.H{"Title": "Наши Мастера"})
}

func AuthPage(c *gin.Context) {
	c.HTML(200, "auth.html", gin.H{"Title": "Вход"})
}

func ContactsPage(c *gin.Context) {
	c.HTML(200, "contacts.html", gin.H{"Title": "Контакты"})
}

func AdminPage(c *gin.Context) {
	c.HTML(200, "admin.html", gin.H{"Title": "Страница Администратора"})
}

func UserPage(c *gin.Context) {
	c.HTML(200, "user.html", gin.H{"Title": "Страница Юзера"})
}

func AdminPanelPage(c *gin.Context) {
	c.HTML(200, "admin_panel.html", gin.H{"Title": "Страница Админа"})
}

func AppointmentPage(c *gin.Context) {
	c.HTML(200, "appointment.html", gin.H{"Title": "Страница записи"})
}
