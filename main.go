// `main.go`
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create Gin engine with Logger and Recovery middleware
	r := gin.Default()

	// Add routes here
	// ...

	// Run the server
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
