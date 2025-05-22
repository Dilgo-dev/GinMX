package handlers

import (
	"html/template"
	"net/http"

	"github.com/Dilgo-dev/GinMX/internal/config"
	"github.com/Dilgo-dev/GinMX/internal/models"
	"github.com/Dilgo-dev/GinMX/internal/utils"
	"github.com/gin-gonic/gin"
)

func GetBlog(c *gin.Context) {
	var blogs []models.Blog
	config.GetDB().Find(&blogs)

	// Rendu du template "home" dans la base
	homeContent, _ := utils.RenderTemplate("home", gin.H{"blogs": blogs})

	c.HTML(http.StatusOK, "base", gin.H{
		"title":   "Gin + HTMX 🦥",
		"content": template.HTML(homeContent),
	})
}
