package sqlstorage

import (
	"context"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/model"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	Db *sqlx.DB
}

func New(db *sqlx.DB) *Storage {
	return &Storage{Db: db}
}

// Создание события.
func (r *Storage) CreateEvent(ctx context.Context, event *model.Event) error {
	query := `
				INSERT INTO events (title, event_time, duration, description, user_id, time_to_notify)
				VALUES (:title, :event_time, :duration, :description, :user_id, :time_to_notify)
`

	if _, err := r.Db.NamedExecContext(ctx, query, event); err != nil {
		return err
	}
	return nil
}

// Обновление события.
func (r *Storage) UpdateEvent(ctx context.Context, event *model.Event) error {
	query := `
				UPDATE events
				SET title          = :title,
				    event_time     = :event_time,
				    duration       = :duration,
				    description    = :description,
				    user_id        = :user_id,
				    time_to_notify = :time_to_notify
				WHERE id = :id;
 `
	if _, err := r.Db.NamedExecContext(ctx, query, event); err != nil {
		return err
	}
	return nil
}

// Получения события по ID.
func (r *Storage) GetEventByID(ctx context.Context, id string) (*model.Event, error) {
	event := &model.Event{}
	err := r.Db.GetContext(ctx, event, "SELECT * FROM events WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return event, nil
}
