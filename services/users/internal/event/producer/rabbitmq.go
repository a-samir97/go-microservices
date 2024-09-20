package producer

import (
	"encoding/json"
	"log"
	"users/internal/event"

	ampq "github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher struct {
	// rabbitMQ sdk
	Writer    ampq.Channel
	QueueName string
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

	q, err := ch.QueueDeclare("users", false, false, false, false, nil)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return &RabbitMQPublisher{Writer: *ch, QueueName: q.Name}
}

func (rp *RabbitMQPublisher) PublishUserCreatedEvent(userCreatedEvent event.UserCreatedEvent) error {
	value, err := json.Marshal(userCreatedEvent)

	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	err = rp.Writer.Publish(
		"",
		rp.QueueName,
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

func (rp *RabbitMQPublisher) PublishUserUpdatedEvent(userUpdatedEvent event.UserUpdatedEvent) error {
	value, err := json.Marshal(userUpdatedEvent)

	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	err = rp.Writer.Publish(
		"",
		rp.QueueName,
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

func (rp *RabbitMQPublisher) PublishUserDeletedEvent(userDeletedEvent event.UserDeletedEvent) error {
	value, err := json.Marshal(userDeletedEvent)

	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	err = rp.Writer.Publish(
		"",
		rp.QueueName,
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
