package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Dilgo-dev/GinMX/internal/config"
	"github.com/Dilgo-dev/GinMX/internal/models"
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

	r.GET("/", func(c *gin.Context) {
		var blogs []models.Blog
		config.GetDB().Find(&blogs)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Gin + HTMX 🦥",
			"blogs": blogs,
		})
	})

	r.Run(":" + port)
}