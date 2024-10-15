package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	PORT = 8080
)

func main() {
	router := gin.Default()

	router.GET("/hi", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})

	router.Run(fmt.Sprintf(":%d", PORT))
}
