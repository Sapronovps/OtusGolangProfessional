package app

import (
	"context"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/model"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/storage"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	_ "github.com/lib/pq" // for postgres
)

type App struct {
	logger  *zap.Logger
	storage storage.Storage
}

func New(logger *zap.Logger, storage storage.Storage) *App {
	return &App{logger: logger, storage: storage}
}

func NewDB(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (a *App) CreateEvent(_ context.Context, title string) error {
	event := &model.Event{Title: title}
	return a.storage.Event().Create(event)
}
