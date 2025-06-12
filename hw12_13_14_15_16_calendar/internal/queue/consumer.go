package queue

import "github.com/rabbitmq/amqp091-go"

type Consumer struct {
	dsn          string
	QueueName    string
	exchangeName string
	routingKey   string
	conn         *amqp091.Connection
	channel      *amqp091.Channel
}

func NewConsumer(dsn, queueName, exchangeName, routingKey string) *Consumer {
	return &Consumer{dsn: dsn, QueueName: queueName, exchangeName: exchangeName, routingKey: routingKey}
}

func (c *Consumer) Consume() *amqp091.Channel {
	if c.conn == nil {
		c.conn = createConnection(c.dsn)
	}
	if c.channel == nil {
		c.channel = createChannel(c.conn)
		createExchange(c.channel, c.exchangeName)
		createQueue(c.channel, c.QueueName)
		queueBind(c.channel, c.QueueName, c.routingKey, c.exchangeName)
	}

	return c.channel
}

func createQueue(channel *amqp091.Channel, queueName string) *amqp091.Queue {
	queue, err := channel.QueueDeclare(
		queueName,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // args
	)
	failOnError(err, "Failed to declare a queue")

	return &queue
}

func queueBind(channel *amqp091.Channel, queueName, routingKey, exchangeName string) {
	err := channel.QueueBind(
		queueName,
		routingKey,
		exchangeName,
		false, // no-wait
		nil,   // args
	)
	failOnError(err, "Failed to bind a queue")
}
