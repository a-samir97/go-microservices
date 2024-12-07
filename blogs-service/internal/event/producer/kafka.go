package producer

import (
	"blogs/internal/event"
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
	kafka.DialLeader(context.Background(), "tcp", "kafka:9092", "blog_likes", 0)
	kafka.DialLeader(context.Background(), "tcp", "kafka:9092", "blog_dislikes", 0)
	kafka.DialLeader(context.Background(), "tcp", "kafka:9092", "blog_claps", 0)
	kafka.DialLeader(context.Background(), "tcp", "kafka:9092", "blog_created", 0)

	if err != nil {
		panic(err.Error())
	}
	w := kafka.Writer{
		Addr:     kafka.TCP("kafka:9092"),
		Balancer: &kafka.LeastBytes{}, // load balancer
	}

	return &KafkaPublisher{Writer: &w}
}

func (kp *KafkaPublisher) PublishBlogCreated(blogEvent event.BlogCreatedEvent) {
	value, err := json.Marshal(blogEvent)

	if err != nil {
		log.Println(err.Error())
	}
	err = kp.Writer.WriteMessages(context.Background(), kafka.Message{Value: value, Topic: "blog_created"})

	if err != nil {
		log.Println(err.Error())
	}
}

func (kp *KafkaPublisher) PublishBlogViewed(blogEvent event.BlogViewedEvent) {
	value, err := json.Marshal(blogEvent)

	if err != nil {
		log.Println(err.Error())
	}
	err = kp.Writer.WriteMessages(context.Background(), kafka.Message{Value: value, Topic: "blog_viewed"})

	if err != nil {
		log.Println(err.Error())
	}
}

func (kp *KafkaPublisher) PublishBlogLiked(blogLiked event.BlogLikedEvent) {
	value, err := json.Marshal(blogLiked)

	if err != nil {
		log.Println(err.Error())
	}
	err = kp.Writer.WriteMessages(context.Background(), kafka.Message{Value: value, Topic: "blog_likes"})

	if err != nil {
		log.Println(err.Error())
	}
}

func (kp *KafkaPublisher) PublishBlogDisliked(blogDisliked event.BlogDislikeEvent) {
	value, err := json.Marshal(blogDisliked)

	if err != nil {
		log.Println(err.Error())
	}
	err = kp.Writer.WriteMessages(context.Background(), kafka.Message{Value: value, Topic: "blog_dislikes"})

	if err != nil {
		log.Println(err.Error())
	}
}

func (kp *KafkaPublisher) PublishBlogClaped(blogClapped event.BlogClappedEvent) {
	value, err := json.Marshal(blogClapped)

	if err != nil {
		log.Println(err.Error())
	}
	err = kp.Writer.WriteMessages(context.Background(), kafka.Message{Value: value, Topic: "blog_claps"})

	if err != nil {
		log.Println(err.Error())
	}
}
