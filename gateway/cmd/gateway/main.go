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

func getPostsHandler(c *gin.Context) {
	// https://blog.logrocket.com/making-http-requests-in-go/
	// Get posts
	url := "https://jsonplaceholder.typicode.com/posts"
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Error fetching posts from resource.",
			"error":   err,
		})
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error reading response body.",
			"error":   err,
		})
		return
	}

	// Parse into JSON structure
	var result []map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to parse JSON body",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func getPostHandler(c *gin.Context) {
	postId := c.Param("id")
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s", postId)
	resp, err := http.Get(url)
	fmt.Println(resp.StatusCode)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println(err)
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Error fetching post from resource",
			"error":   err,
		})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error reading response body.",
			"error":   err,
		})
		return
	}

	// Parse into JSON structure
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to parse JSON body",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func main() {
	router := gin.Default()

	router.GET("/posts", getPostsHandler)
	router.GET("/posts/:id", getPostHandler)

	router.GET("/hi", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})

	router.Run(fmt.Sprintf(":%d", PORT))
}
