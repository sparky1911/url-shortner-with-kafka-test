package repository

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

type KafkaRepository struct {
	writer *kafka.Writer
}

func NewKafkaRepository(broker string) *KafkaRepository {
	return &KafkaRepository{
		writer: &kafka.Writer{
			Addr:                   kafka.TCP(broker),
			Topic:                  "url_clicks",
			Balancer:               &kafka.LeastBytes{},
			AllowAutoTopicCreation: true,
		},
	}
}


type ClickEvent struct{ 
	Code string `josn:"code"`
	IP string `josn:"ip"`
	UserAgent string `josn:"user_agent"`
}

func (k *KafkaRepository) PublishClick(ctx context.Context, code,ip,userAgent string ) error {
	event:=&ClickEvent{
		Code: code,
		IP:ip,
		UserAgent: userAgent,
	}
	payload,_:=json.Marshal(event)

	return k.writer.WriteMessages(ctx,kafka.Message{Value: payload})

}
