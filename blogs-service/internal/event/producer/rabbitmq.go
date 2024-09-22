package producer

import (
	"encoding/json"
	"log"

	ampq "github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher struct {
	// rabbitMQ sdk
	Writer *ampq.Channel
}

func NewRabbitMQPublisher() *RabbitMQPublisher {
	conn, err := ampq.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		log.Fatalln(err.Error())
	}

	ch, err := conn.Channel()

	if err != nil {
		log.Fatalln(err.Error())
	}

	return &RabbitMQPublisher{Writer: ch}
}

func (rp *RabbitMQPublisher) PublishEvent(blogEvent interface{}, queue string) error {
	value, err := json.Marshal(blogEvent)

	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	err = rp.Writer.Publish(
		"",
		queue,
		false,
		false,
		ampq.Publishing{
			Body: value,
		},
	)

	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	return nil
}
