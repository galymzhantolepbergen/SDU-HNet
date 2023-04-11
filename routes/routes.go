// routes.go
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*")

	// Route for home page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home_page.html", gin.H{})
	})
}
