package handlers

import (
	"net/http"

	"github.com/Dilgo-dev/GinMX/internal/config"
	"github.com/Dilgo-dev/GinMX/internal/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Gin + HTMX - Register 🦥",
	})
}

func RegisterPost(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "register.html", gin.H{
			"title": "Gin + HTMX - Register 🦥",
			"error": "Failed to hash password",
		})
	}

	user := models.User{Email: email, Password: string(hashedPassword)}
	config.GetDB().Create(&user)

	c.Redirect(http.StatusFound, "/")
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Gin + HTMX - Login 🦥",
	})
}

func LoginPost(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	var user models.User
	result := config.GetDB().Where("email = ?", email).First(&user)
	if result.Error != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"title": "Gin + HTMX - Login 🦥",
			"error": "Invalid email or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"title": "Gin + HTMX - Login 🦥",
			"error": "Invalid email or password",
		})
		return
	}

	c.Redirect(http.StatusFound, "/")
}