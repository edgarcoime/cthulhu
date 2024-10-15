package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	rabbit "github.com/edgarcoime/cthulhu/internal/app/rabbit"
	"github.com/gin-gonic/gin"
)

type PostRabbitMessageRequest struct {
	Message string `json:"message" binding:"required"`
}

func PostRabbitMessage(c *gin.Context) {
	// Grab rmq service
	rmq, exists := c.MustGet("rmq").(*rabbit.RabbitMQService)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "RabbitMQ service not found.",
			"error":   nil,
		})
		return
	}

	var msgReq PostRabbitMessageRequest
	if err := c.ShouldBindJSON(&msgReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request, please provide a valid 'message' field.",
			"error":   err,
		})
		return
	}

	err := rmq.SendMessage(rabbit.QUEUE_NAME, msgReq.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error Could not send message to RabbitMQ",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": msgReq.Message,
	})
}

func GetPostsHandler(c *gin.Context) {
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

func GetPostHandler(c *gin.Context) {
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
