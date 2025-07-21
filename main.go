// server/main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong from server",
		})
	})

	router.GET("/greet", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			name = "Guest"
		}
		c.JSON(http.StatusOK, gin.H{
			"greeting": "Hello " + name,
		})
	})

	router.Run(":8081") // start server on port 8081
}
