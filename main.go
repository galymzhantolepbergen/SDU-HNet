package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()
	println("http://localhost:8080/")

	// Switch to release mode
	gin.SetMode(gin.ReleaseMode)

	// Define a route and a handler function
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// Start the server
	router.Run(":8080")
}
