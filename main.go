package main

import (
	"barbershop/config"
	"barbershop/routes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("Request - Method: %s | Status: %d | Duration: %v", c.Request.Method, c.Writer.Status(), duration)
	}
}

func ProtectionPages() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/css" || c.Request.URL.Path == "/img" {
			// Возвращаем ошибку 404
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.Next()
	}
}

func main() {
	config.InitDB()
	r := routes.SetupRouter()
	r.Static("/static", "./static")
	r.Static("/css", "./static/css")
	r.Static("/js", "./static/js")
	r.Static("/img", "./static/img")
	r.LoadHTMLGlob("static/html/*")
	// r.LoadHTMLGlob("static/*.html")
	fmt.Println("Server is running on http://localhost:8080")
	r.Use(LoggerMiddleware(), ProtectionPages())
	r.Run(":8080")

}
