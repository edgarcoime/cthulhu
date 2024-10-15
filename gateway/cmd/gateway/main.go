package main

import (
	"fmt"
	"net/http"

	rabbit "github.com/edgarcoime/cthulhu/internal/app/rabbit"
	routes "github.com/edgarcoime/cthulhu/internal/app/routes"
	"github.com/edgarcoime/cthulhu/internal/pkg"
	"github.com/gin-gonic/gin"
)

// Setup gateway Struct that has the router and rmqservice with open connections
// Can either pass rmq as middleware or use closure
func WithRMQ(rmq *rabbit.RabbitMQService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Channels should be created and closed per request as they are not thread safe
		// Connections can be reused but queues and channels should be created as per connection
		// These systems are lightweight and can be built and torn down easily
		c.Set("rmq", rmq)
		c.Next()
	}
}

func main() {
	// https://www.youtube.com/watch?v=bfVddTJNiAw
	// https://www.rabbitmq.com/tutorials/tutorial-one-go
	// https://www.svix.com/resources/guides/rabbitmq-docker-setup-guide
	fmt.Println("Running RabbitMQ service...")
	rmq := rabbit.NewRabbitMQService()
	defer rmq.Close()

	// Create a queue
	qName := rabbit.QUEUE_NAME
	q, err := rmq.Ch.QueueDeclare(
		qName, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	rabbit.FailOnError(err, "Failed to declare a queue.")
	rmq.Queues[qName] = &q

	// Setup http service
	router := gin.Default()
	router.Use(WithRMQ(rmq))
	router.GET("/posts", routes.GetPostsHandler)
	router.GET("/posts/:id", routes.GetPostHandler)
	router.POST("/rabbit/msg", routes.PostRabbitMessage)

	router.GET("/hi", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin!",
		})
	})

	router.Run(fmt.Sprintf(":%d", pkg.HTTP_PORT))
}
