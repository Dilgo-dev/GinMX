package handlers

import (
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/Dilgo-dev/GinMX/internal/config"
	"github.com/Dilgo-dev/GinMX/internal/models"
	"github.com/Dilgo-dev/GinMX/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterPage(c *gin.Context) {
	registerContent, _ := utils.RenderTemplate("register", gin.H{})

	c.HTML(http.StatusOK, "base", gin.H{
		"title":   "Gin + HTMX - Register 🦥",
		"content": template.HTML(registerContent),
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
	loginContent, _ := utils.RenderTemplate("login", gin.H{})

	c.HTML(http.StatusOK, "base", gin.H{
		"title":   "Gin + HTMX - Login 🦥",
		"content": template.HTML(loginContent),
	})
}

func LoginPost(c *gin.Context) {
	if os.Getenv("SECRET_KEY") == "" {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"title": "Gin + HTMX - Login 🦥",
			"error": "SECRET_KEY is not set",
		})
		return
	}

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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": user.Email,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"title": "Gin + HTMX - Login 🦥",
			"error": "Failed to generate token",
		})
	}

	c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/")
}
