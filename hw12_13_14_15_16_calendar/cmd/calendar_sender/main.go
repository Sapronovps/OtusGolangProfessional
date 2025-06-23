package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/logger"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/queue"
	sender2 "github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/sender"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "/etc/calendar-sender/config.yaml", "Path to configuration file")
	flag.Parse()

	config := NewConfig(configFile)
	logg := logger.New(config.Logger.Level, config.Logger.File)

	dsn := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		config.RabbitMQ.User, config.RabbitMQ.Password, config.RabbitMQ.Host, config.RabbitMQ.Port,
	)
	consumer := queue.NewConsumer(
		dsn,
		config.RabbitMQ.QueueName,
		config.RabbitMQ.ExchangeName,
		config.RabbitMQ.RoutingKey)

	sender := sender2.New(logg, *consumer)

	// Контекст для публикации сообщения
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sender.Run(ctx)
}
