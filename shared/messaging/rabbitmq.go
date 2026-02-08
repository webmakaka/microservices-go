package messaging

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn *amqp.Connection
}

func NewRabbitMQ(uri string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, fmt.Errorf("[Me] Failed to connect to RabbitMQ: %v", err)
	}

	rmq := &RabbitMQ{
		conn: conn,
	}
	return rmq, nil
}

func (r *RabbitMQ) Close() {
	if err := r.conn.Close(); err != nil {
		return
	}
}
