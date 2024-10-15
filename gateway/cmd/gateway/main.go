package main

import (
	"fmt"
	"net/http"

	rabbitmq "github.com/edgarcoime/cthulhu/internal/app/rabbit"
	routes "github.com/edgarcoime/cthulhu/internal/app/routes"
	"github.com/edgarcoime/cthulhu/internal/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	rabbitmq.Run()

	// Setup http service
	router := gin.Default()
	router.GET("/posts", routes.GetPostsHandler)
	router.GET("/posts/:id", routes.GetPostHandler)

	router.GET("/hi", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})

	router.Run(fmt.Sprintf(":%d", pkg.HTTP_PORT))
}
