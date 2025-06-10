package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/app"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/logger"
	sqlstorage "github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/storage/sql"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "config", "/etc/calendar-scheduler/config.yaml", "Path to configuration file")
	flag.Parse()

	config := NewConfig(configFile)
	logg := logger.New(config.Logger.Level, config.Logger.File)

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Username, config.DB.Password, config.DB.DBName,
	)
	db, err := app.NewDB(dsn)
	if err != nil {
		panic(fmt.Errorf("failed connect to db: %w", err))
	}
	storage := sqlstorage.New(db)

	calendar := app.New(logg, storage)
	timeMonth, _ := time.Parse("2006-01-02 15:04:05", "2025-06-01 00:00:00")
	events, err := calendar.ListByMonth(timeMonth)
	if err != nil {
		panic(fmt.Errorf("failed list events: %w", err))
	}

	fmt.Println(events)

	// Код, который уйдет в internal
	// Подключение к RabbitMQ
	dsn = fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		config.RabbitMQ.User, config.RabbitMQ.Password, config.RabbitMQ.Host, config.RabbitMQ.Port,
	)
	conn, err := amqp.Dial(dsn)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Создание канала
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Создание exchange (если не существует)
	exchangeName := "my_topic_exchange"
	err = ch.ExchangeDeclare(
		exchangeName, // имя exchange
		"topic",      // тип exchange (topic позволяет использовать routing keys с масками)
		true,         // durable (переживет перезагрузку сервера)
		false,        // auto-deleted (удаляется, когда нте подписчиков)
		false,        // internal (только для внутреннего использования)
		false,        // no-wait (не ждать подтверждения)
		nil,          // аргументы
	)
	failOnError(err, "Failed to declare an exchange")

	// Контекст для публикации сообщения
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Пример routing key (может быть любой строкой, например user.created)
	routingKey := "example.key"

	// Тело сообщения
	body := fmt.Sprintf("Hello RabbitMQ! Time: %v", time.Now())

	// Публикация сообщения
	err = ch.PublishWithContext(
		ctx,
		exchangeName, //exchange
		routingKey,   // routing key
		false,        // mandatory (если true, сообщение вернется, если не доставлено)
		false,        // immediate (если true, сообщение вернется, если не доставлено немедленно)
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	failOnError(err, "Failed to publish a message")

	logg.Info(fmt.Sprintf(" [x] Sent %s with routing key '%s'", body, routingKey))
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
