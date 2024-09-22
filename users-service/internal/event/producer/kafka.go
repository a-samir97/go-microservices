package producer

import (
	"context"
	"encoding/json"
	"users/internal/event"

	"github.com/segmentio/kafka-go"
)

type KafkaPublisher struct {
	Writer kafka.Writer
}

func NewKafkaPublisher() *KafkaPublisher {
	w := kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "users",
		Balancer: &kafka.LeastBytes{}, // load balancer
	}

	return &KafkaPublisher{Writer: w}
}

func (kp *KafkaPublisher) PublishUserCreatedEvent(userCreatedEvent event.UserCreatedEvent) error {
	value, err := json.Marshal(userCreatedEvent)

	if err != nil {
		return err
	}
	err = kp.Writer.WriteMessages(context.Background(), kafka.Message{Value: value})

	if err != nil {
		return err
	}

	return nil
}

func (kp *KafkaPublisher) PublishUserUpdatedEvent(userUpdatedEvent event.UserUpdatedEvent) error {
	value, err := json.Marshal(userUpdatedEvent)

	if err != nil {
		return err
	}

	err = kp.Writer.WriteMessages(context.Background(), kafka.Message{Value: value})

	if err != nil {
		return err
	}
	return nil
}

func (kp *KafkaPublisher) PublishUserDeletedEvent(userDeletedEvent event.UserDeletedEvent) error {
	value, err := json.Marshal(userDeletedEvent)

	if err != nil {
		return err
	}

	err = kp.Writer.WriteMessages(context.Background(), kafka.Message{Value: value})

	if err != nil {
		return err
	}

	return nil
}
