package app

import (
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/model"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/storage"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"time"

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

func (a *App) GetEvent(id int) (model.Event, error) {
	return a.storage.Event().Get(id)
}

func (a *App) UpdateEvent(id int, event *model.Event) error {
	return a.storage.Event().Update(id, event)
}

func (a *App) CreateEvent(event *model.Event) error {
	return a.storage.Event().Create(event)
}

func (a *App) ListByDay(date time.Time) ([]model.Event, error) {
	return a.storage.Event().ListByDay(date)
}

func (a *App) ListByWeek(date time.Time) ([]model.Event, error) {
	return a.storage.Event().ListByWeek(date)
}

func (a *App) ListByMonth(date time.Time) ([]model.Event, error) {
	return a.storage.Event().ListByMonth(date)
}

func (a *App) DeleteEvent(id int) error {
	return a.storage.Event().Delete(id)
}
