package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/app"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/logger"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/queue"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/scheduler"
	sqlstorage "github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/storage/sql"
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

	dsn = fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		config.RabbitMQ.User, config.RabbitMQ.Password, config.RabbitMQ.Host, config.RabbitMQ.Port,
	)
	producer := queue.New(dsn, config.RabbitMQ.ExchangeName, config.RabbitMQ.RoutingKey)

	schedule := scheduler.New(logg, storage, *producer, config.RabbitMQ.ScanInterval)
	// Контекст для публикации сообщения
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	schedule.Run(ctx)
}
