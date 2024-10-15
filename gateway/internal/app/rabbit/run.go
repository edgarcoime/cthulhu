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

func failOnError(err error, msg string) {
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
	failOnError(err, "Failed to connect to RabbitMQ.")

	// Create channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel.")

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
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
	return nil
}

func Run() {
	fmt.Println("Running RabbitMQ service...")
	rmq := NewRabbitMQService()
	defer rmq.Close()

	// Create a queue
	qName := QUEUE_NAME
	q, err := rmq.Ch.QueueDeclare(
		qName, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue.")
	rmq.Queues[qName] = &q

	// send Message
	rmq.SendMessage(qName, "Hello RabbitMQ!")
}
