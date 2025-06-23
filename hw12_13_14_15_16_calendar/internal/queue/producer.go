package queue

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

type Producer struct {
	dsn          string
	exchangeName string
	routingKey   string
	Conn         *amqp091.Connection
	Channel      *amqp091.Channel
}

func New(dsn string, exchangeName string, routingKey string) *Producer {
	return &Producer{dsn: dsn, exchangeName: exchangeName, routingKey: routingKey}
}

func (p *Producer) PublishWithContext(ctx context.Context, msg []byte) error {
	if p.Conn == nil {
		p.Conn = createConnection(p.dsn)
	}
	if p.Channel == nil {
		p.Channel = createChannel(p.Conn)
		createExchange(p.Channel, p.exchangeName)
	}
	return publish(ctx, p.Channel, p.exchangeName, p.routingKey, msg)
}

func createConnection(dsn string) *amqp091.Connection {
	conn, err := amqp091.Dial(dsn)
	failOnError(err, "Failed to connect to RabbitMQ")

	return conn
}

func createChannel(conn *amqp091.Connection) *amqp091.Channel {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	return ch
}

func createExchange(channel *amqp091.Channel, exchangeName string) {
	err := channel.ExchangeDeclare(
		exchangeName, // имя exchange
		"topic",      // тип exchange (topic позволяет использовать routing keys с масками)
		true,         // durable (переживет перезагрузку сервера)
		false,        // auto-deleted (удаляется, когда нет подписчиков)
		false,        // internal (только для внутреннего использования)
		false,        // no-wait (не ждать подтверждение)
		nil,          // аргументы
	)
	failOnError(err, "Failed to declare an exchange")
}

func publish(ctx context.Context, channel *amqp091.Channel, exchangeName string, routingKey string, msg []byte) error {
	return channel.PublishWithContext(
		ctx,
		exchangeName,
		routingKey,
		false, // mandatory (если true, сообщение вернется, если не доставлено)
		false, // immediate (если true, сообщение вернется, если не доставлено немедленно)
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
