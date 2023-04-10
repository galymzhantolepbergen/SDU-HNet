package main

import (
	"github.com/galymzhantolepbergen/SDU-HNet/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	println("http://localhost:8080/")

	// Set up routes
	routes.SetupRoutes(router)

	// Start the server
	router.Run(":8080")
}
