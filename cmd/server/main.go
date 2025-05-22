package main

import (
	"log"
	"os"

	"github.com/Dilgo-dev/GinMX/internal/config"
	"github.com/Dilgo-dev/GinMX/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	config.InitDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := gin.Default()
	r.LoadHTMLGlob("internal/templates/*")
	r.Static("/static", "./static")

	r.GET("/", handlers.GetBlog)
	r.GET("/register", handlers.RegisterPage)
	r.POST("/register", handlers.RegisterPost)
	r.GET("/login", handlers.LoginPage)
	r.POST("/login", handlers.LoginPost)

	r.Run(":" + port)
}