package kafka

import (
	"context"
	"log"

	kafkago "github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafkago.Reader
}

func NewConsumer(broker string, topics []string, groupID string) *Consumer {
	return &Consumer{
		reader: kafkago.NewReader(kafkago.ReaderConfig{
			Brokers:     []string{broker},
			GroupID:     groupID,
			GroupTopics: topics,
			MinBytes:    1,
			MaxBytes:    10e6,
		}),
	}
}

// Start memblokir sampai ctx dibatalkan. handle dipanggil per-pesan dari topic mana pun.
func (c *Consumer) Start(ctx context.Context, handle func(topic string, data []byte) error) {
	for {
		msg, err := c.reader.FetchMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			log.Println("kafka fetch error:", err)
			continue
		}
		if err := handle(msg.Topic, msg.Value); err != nil {
			log.Printf("kafka handle error topic=%s: %v", msg.Topic, err)
			continue
		}
		if err := c.reader.CommitMessages(ctx, msg); err != nil {
			log.Println("kafka commit error:", err)
		}
	}
}
