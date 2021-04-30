package listener

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/telepenin/website-checker/shared"
)

type KafkaListener struct {
	Config shared.Kafka
	Conn   *kafka.Conn
}

func Init(config shared.Kafka) (*KafkaListener, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp",
		config.Address, config.Topic, config.Partition)
	if err != nil {
		return nil, fmt.Errorf("failed to dial leader: %w", err)
	}

	return &KafkaListener{
		Config: config,
		Conn:   conn,
	}, nil
}

func (l *KafkaListener) Close() error {
	return l.Conn.Close()
}

func (l *KafkaListener) Process(resp *shared.Response) error {
	b, err := resp.ToJson()
	if err != nil {
		return err
	}

	_, err = l.Conn.WriteMessages(
		kafka.Message{Value: b},
	)
	if err != nil {
		return err
	}
	return nil

}
