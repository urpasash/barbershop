package routes

import (
	"barbershop/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	// config.DB.AutoMigrate(&models.User{}, &models.Master{}, &models.Appointment{})

	r.GET("/", controllers.HomePage)
	r.GET("/services", controllers.ServicesPage)
	r.GET("/masters", controllers.MastersPage)
	r.GET("/contacts", controllers.ContactsPage)
	r.GET("/admin", controllers.AdminPage)
	r.GET("/auth", controllers.AuthPage)
	r.GET("/user", controllers.UserPage)
	r.GET("/admin_panel", controllers.AdminPanelPage)
	r.GET("/appointment", controllers.AppointmentPage)

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.AddUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.POST("/appointment", controllers.CreateAppointment)

	r.GET("/appointments", controllers.GetAppointment)
	r.PUT("/appointments/:id", controllers.UpdateAppointment)
	r.DELETE("/appointments/:id", controllers.DeleteAppointment)

	// r.GET("/masters", controllers.GetMasters)

	// r.GET("/records", controllers.GetRecords)
	// r.POST("/records", controllers.AddRecord)
	// r.PUT("/records/:id", controllers.UpdateRecord)
	// r.DELETE("/records/:id", controllers.DeleteRecord)
	// r.GET("/set-cookie", setCookieHandler)
	// r.GET("/read-cookie", readCookieHandler)
	// r.GET("/delete-cookie", deleteCookieHandler)

	return r
}
