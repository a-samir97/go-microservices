package producer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaPublisher struct {
	Writer *kafka.Writer
}

func NewKafkaPublisher() *KafkaPublisher {
	_, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", "blog_viewed", 0)
	if err != nil {
		panic(err.Error())
	}
	w := kafka.Writer{
		Addr:     kafka.TCP("kafka:9092"),
		Topic:    "blog_viewed",
		Balancer: &kafka.LeastBytes{}, // load balancer
	}

	return &KafkaPublisher{Writer: &w}
}

func (kp *KafkaPublisher) PublishEvent(blogEvent interface{}, topic string) error {
	value, err := json.Marshal(blogEvent)

	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = kp.Writer.WriteMessages(context.Background(), kafka.Message{Value: value})

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
