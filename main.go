package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {
		// Allowed origins (Angular frontend + Backend)
		allowedOrigins := []string{
			"http://52.67.71.173:4200", // Angular dev server
			"http://52.67.71.173:8080", // Your Gin backend (if needed)
			"http://localhost:4200",    // Localhost for Angular dev server
		}

		// Get the request's Origin
		requestOrigin := c.Request.Header.Get("Origin")

		// Check if the origin is allowed
		for _, origin := range allowedOrigins {
			if requestOrigin == origin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", requestOrigin)
				break
			}
		}

		// Required headers
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle OPTIONS (preflight)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	// Function call routes
	setupRoutes(r)

	// Corre en puerto 8080
	r.Run()

}
