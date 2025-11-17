package publishers

import (
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"

	"shared"
)

type RabbitMQEventPublisher struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewRabbitMQEventPublisher() *RabbitMQEventPublisher {
	connection, error := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if error != nil {
		log.Panicf("Failed to connect to RabbitMQ: %v", error)
	}

	channel, error := connection.Channel()

	if error != nil {
		log.Panicf("Failed to create channel: %v", error)
	}

	return &RabbitMQEventPublisher{
		connection: connection,
		channel:    channel,
	}
}

func (publisher *RabbitMQEventPublisher) Dispose() {
	defer publisher.channel.Close()
	defer publisher.connection.Close()
}

func (publisher *RabbitMQEventPublisher) Publish(event shared.Event) error {
	queue, error := publisher.channel.QueueDeclare(
		event.GetType(),
		false,
		false,
		false,
		false,
		nil,
	)

	if error != nil {
		log.Printf("Failed to declare queue: %v", error)
		return error
	}

	body, error := json.Marshal(event)

	if error != nil {
		log.Printf("Failed to marshal event: %v", error)
		return error
	}

	error = publisher.channel.PublishWithContext(
		context.Background(),
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if error != nil {
		log.Printf("Failed to publish event: %v", error)
		return error
	}

	return nil
}
