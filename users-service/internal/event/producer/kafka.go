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
	_, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", "userCreated", 0)
	kafka.DialLeader(context.Background(), "tcp", "kafka:9092", "userDeleted", 0)
	kafka.DialLeader(context.Background(), "tcp", "kafka:9092", "userUpdated", 0)

	if err != nil {
		panic(err.Error())
	}
	w := kafka.Writer{
		Addr:     kafka.TCP("kafka:9092"),
		Balancer: &kafka.LeastBytes{}, // load balancer
	}

	return &KafkaPublisher{Writer: w}
}

func (kp *KafkaPublisher) PublishUserCreatedEvent(userCreatedEvent event.UserCreatedEvent) error {
	value, err := json.Marshal(userCreatedEvent)

	if err != nil {
		return err
	}
	err = kp.Writer.WriteMessages(context.Background(), kafka.Message{Value: value, Topic: "userCreated"})

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

	err = kp.Writer.WriteMessages(context.Background(), kafka.Message{Value: value, Topic: "userUpdated"})

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

	err = kp.Writer.WriteMessages(context.Background(), kafka.Message{Value: value, Topic: "userDeleted"})

	if err != nil {
		return err
	}

	return nil
}
