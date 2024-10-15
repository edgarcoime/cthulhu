package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	PORT = 8080
)

func main() {
	router := gin.Default()

	router.GET("/posts", func(c *gin.Context) {
		// https://blog.logrocket.com/making-http-requests-in-go/
		// Get posts
		url := "https://jsonplaceholder.typicode.com/posts"
		resp, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"message": "Error fetching posts from resource.",
			})
			return
		}
		defer resp.Body.Close()

		// Read response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error reading response body.",
			})
			return
		}

		// Parse into JSON structure
		var result []map[string]interface{}
		err = json.Unmarshal(body, &result)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to parse JSON body",
			})
			return
		}

		c.JSON(http.StatusOK, result)
	})

	router.GET("/hi", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})

	router.Run(fmt.Sprintf(":%d", PORT))
}
