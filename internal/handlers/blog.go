package handlers

import (
	"net/http"

	"github.com/Dilgo-dev/GinMX/internal/config"
	"github.com/Dilgo-dev/GinMX/internal/models"
	"github.com/gin-gonic/gin"
)

func GetBlog(c *gin.Context) {
	var blogs []models.Blog
	config.GetDB().Find(&blogs)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Gin + HTMX 🦥",
		"blogs": blogs,
	})
}