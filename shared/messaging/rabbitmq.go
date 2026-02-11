package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"ride-sharing/shared/contracts"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	TripExchange = "trip"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQ(uri string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, fmt.Errorf("[Me] Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("[Me] Failed to create channel: %v", err)
	}

	rmq := &RabbitMQ{
		conn:    conn,
		Channel: ch,
	}

	if err := rmq.setupExchangesAndQueues(); err != nil {
		rmq.Close()
		return nil, fmt.Errorf("[Me] Failed to setup exchanges and queues: %v", err)
	}

	return rmq, nil
}

type MessageHandler func(context.Context, amqp.Delivery) error

func (r *RabbitMQ) ConsumeMessage(queueName string, handler MessageHandler) error {

	err := r.Channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return fmt.Errorf("[Me] Failed to set Qos: %v", err)
	}

	msg, err := r.Channel.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)

	if err != nil {
		return err
	}

	ctx := context.Background()

	go func() {
		for msg := range msg {
			log.Printf("[Me] Received a message: %s", msg.Body)

			if err := handler(ctx, msg); err != nil {
				log.Printf("[Me] [ERROR:] Failed to handle message: %v. Message body: %s", err, msg.Body)

				if nackErr := msg.Nack(false, false); nackErr != nil {
					log.Printf("[Me] [ERROR:] Failed to Nack message: %v", nackErr)
				}

				// Continue to the next message
				continue
			}

			// Only Ack if the handler succeeds
			if ackErr := msg.Ack(false); ackErr != nil {
				log.Printf("[Me] [ERROR:] Failed to Ack message: %v. Message body: %s", ackErr, msg.Body)
			}
		}
	}()
	return nil
}

func (r *RabbitMQ) PublishMessage(ctx context.Context, routingKey string, message contracts.AmqpMessage) error {

	log.Printf("[Me] Publishing message with routing key: %s", routingKey)

	jsonMsg, err := json.Marshal(message)

	if err != nil {
		return fmt.Errorf("[Me] Failed to marshal message: %v", err)
	}

	return r.Channel.PublishWithContext(ctx,
		TripExchange, routingKey, false, false, amqp.Publishing{
			ContentType:  "text/plain",
			Body:         jsonMsg,
			DeliveryMode: amqp.Persistent,
		})
}

func (r *RabbitMQ) setupExchangesAndQueues() error {

	err := r.Channel.ExchangeDeclare(
		TripExchange, "topic", true, false, false, false, nil,
	)

	if err != nil {
		return fmt.Errorf("[Me] Failed to declare exchange: %s: %v ", TripExchange, err)
	}

	if err := r.declareAndBindQueue(FindAvailableDriversQueue, []string{contracts.TripEventCreated, contracts.TripEventDriverNotInterested}, TripExchange); err != nil {
		return err
	}

	if err := r.declareAndBindQueue(DriverCmdTripRequestQueue, []string{contracts.DriverCmdTripRequest}, TripExchange); err != nil {
		return err
	}

	if err := r.declareAndBindQueue(DriverTripResponseQueue, []string{contracts.DriverCmdTripAccept, contracts.DriverCmdTripDecline}, TripExchange); err != nil {
		return err
	}

	if err := r.declareAndBindQueue(NotifyDriverNoDriversFoundQueue, []string{contracts.TripEventNoDriversFound}, TripExchange); err != nil {
		return err
	}

	if err := r.declareAndBindQueue(NotifyDriverAssignQueue, []string{contracts.TripEventDriverAssigned}, TripExchange); err != nil {
		return err
	}

	return nil
}

func (r *RabbitMQ) declareAndBindQueue(queueName string, messageTypes []string, exchange string) error {
	q, err := r.Channel.QueueDeclare(
		queueName, true, false, false, false, nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	for _, msg := range messageTypes {
		if err := r.Channel.QueueBind(
			q.Name, msg, exchange, false, nil,
		); err != nil {
			return fmt.Errorf("[Me] Failed to bind queue to %s: %v", queueName, err)
		}
	}

	return nil

}

func (r *RabbitMQ) Close() {
	if r.conn != nil {
		r.conn.Close()
	}

	if r.Channel != nil {
		r.Channel.Close()
	}
}
