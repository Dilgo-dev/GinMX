package utils

import (
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
)

func RenderTemplate(templateName string, data gin.H) (string, error) {
	tmpl, err := template.ParseGlob("internal/templates/*.html")
	if err != nil {
		return "", err
	}

	var buf strings.Builder
	err = tmpl.ExecuteTemplate(&buf, templateName, data)
	return buf.String(), err
}
