package rabbitmq

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/edgarcoime/cthulhu/internal/pkg"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	QUEUE_NAME = "hello"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type RabbitMQService struct {
	Conn   *amqp.Connection
	Ch     *amqp.Channel
	Queues map[string]*amqp.Queue
}

func NewRabbitMQService() *RabbitMQService {
	// Create Connection
	conn, err := amqp.Dial(pkg.RABBITMQ_URL)
	FailOnError(err, "Failed to connect to RabbitMQ.")

	// Create channel
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel.")

	rmq := &RabbitMQService{
		Conn:   conn,
		Ch:     ch,
		Queues: make(map[string]*amqp.Queue),
	}
	return rmq
}

func (rmq *RabbitMQService) Close() {
	defer rmq.Conn.Close()
	defer rmq.Ch.Close()
}

func (rmq *RabbitMQService) SendMessage(qName string, msg string) error {
	// Create Context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Grab queue
	q, ok := rmq.Queues[qName]
	if !ok {
		return errors.New(fmt.Sprintf("Queue with name (%s) does not exist.", qName))
	}

	// Create message
	body := msg
	err := rmq.Ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
	return nil
}

// Either add rmq to middleware so its always within the request
// Create a closure wrapper to have access to rmq
// func WrapRMQWithHandler(rmq *RabbitMQService, handler func(rmq *RabbitMQService, c *gin.Context)) func(c *gin.Context) {
// 	return func(c *gin.Context) {
// 		handler(rmq, c)
// 	}
// }
