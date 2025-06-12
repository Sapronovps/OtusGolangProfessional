package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/queue"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/storage"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Scheduler struct {
	logger       *zap.Logger
	storage      storage.Storage
	producer     queue.Producer
	scanInterval int64
}

func New(logger *zap.Logger, storage storage.Storage, producer queue.Producer, scanInterval int64) *Scheduler {
	return &Scheduler{logger: logger, storage: storage, producer: producer, scanInterval: scanInterval}
}

func (s *Scheduler) Run(ctx context.Context) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Запускаем планировщик
	timeInterval := time.Duration(s.scanInterval) * time.Second
	ticker := time.NewTicker(timeInterval)
	defer ticker.Stop()

	s.logger.Info("Scheduler started. Press Ctrl+C to stop.")

	for {
		select {
		case <-ticker.C:
			events, err := s.storage.Event().ListByMonth(time.Now().AddDate(0, 0, -10))
			if err != nil {
				s.logger.Error("failed to list events", zap.Error(err))
				return
			}

			for _, event := range events {
				eventJSON, err := json.Marshal(event)
				if err != nil {
					s.logger.Error("failed to marshal event", zap.Error(err))
					continue
				}
				err = s.producer.PublishWithContext(ctx, eventJSON)

				if err != nil {
					s.logger.Error("failed to publish queue", zap.Error(err))
					return
				}
				s.logger.Info(fmt.Sprintf(" [x] Sent %s to RabbitMQ", eventJSON))

				// После отправки удаляем событие из БД
			}
		case <-sigChan:
			s.logger.Info("Shutting down gracefully...")
			return
		}
	}
}
