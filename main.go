package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	r.GET("/env", func(c *gin.Context) {
		c.String(200, "Hello, ENV!"+os.Getenv("TEST_ENV"))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err := r.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
