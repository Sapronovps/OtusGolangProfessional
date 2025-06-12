package sender

import (
	"context"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/queue"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

type Sender struct {
	logger   *zap.Logger
	consumer queue.Consumer
}

func New(logger *zap.Logger, consumer queue.Consumer) *Sender {
	return &Sender{logger, consumer}
}

func (s *Sender) Run(ctx context.Context) {
	channel := s.consumer.Consume()

	// Подписываемся на очередь
	msgs, err := channel.Consume(
		s.consumer.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		s.logger.Fatal("failed to consume messages", zap.Error(err))
		return
	}

	sigChang := make(chan os.Signal, 1)
	signal.Notify(sigChang, syscall.SIGINT, syscall.SIGTERM)

	s.logger.Info("Consumer started. Waiting for messages...")

	// Бесконечный цикл обработки сообщений
	for {
		select {
		case msg := <-msgs:
			s.logger.Info("Received message", zap.Any("message", msg.Body))
		case <-sigChang:
			s.logger.Info("Shutting down consumer...")
			return
		case <-ctx.Done():
			return
		}
	}
}
