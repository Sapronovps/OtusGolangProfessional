package sqlstorage

import (
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/model"
	"time"
)

type EventRepository struct {
	storage *Storage
}

func (r *EventRepository) Create(e *model.Event) error {
	query := `
				INSERT INTO events (title, event_time, duration, description, user_id, time_to_notify)
				VALUES (:title, :event_time, :duration, :description, :user_id, :time_to_notify)
`
	if _, err := r.storage.db.NamedExec(query, e); err != nil {
		return err
	}
	return nil
}

func (r *EventRepository) Update(id int, e *model.Event) error {
	e.ID = id
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
	if _, err := r.storage.db.NamedExec(query, e); err != nil {
		return err
	}
	return nil
}

func (r *EventRepository) Delete(id int) error {
	_, err := r.storage.db.Exec("DELETE FROM events WHERE id=$1", id)
	return err
}

func (r *EventRepository) Get(id int) (model.Event, error) {
	event := model.Event{}
	err := r.storage.db.Get(event, "SELECT * FROM events WHERE id = $1", id)
	if err != nil {
		return event, err
	}

	return event, nil
}

func (r *EventRepository) ListByDay(date time.Time) ([]model.Event, error) {
	events := make([]model.Event, 0)
	err := r.storage.db.Select(
		&events,
		`SELECT title,
				       event_time,
				       duration,
				       description,
				       user_id,
				       time_to_notify
				FROM events
				WHERE event_time = $1`,
		date,
	)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r *EventRepository) ListByWeek(date time.Time) ([]model.Event, error) {
	events := make([]model.Event, 0)
	err := r.storage.db.Select(
		&events,
		`SELECT title,
				       event_time,
				       duration,
				       description,
				       user_id,
				       time_to_notify
				FROM events
				WHERE event_time >= $1 and event_time <= event_time + make_interval(days => 7)`,
		date,
	)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r *EventRepository) ListByMonth(date time.Time) ([]model.Event, error) {
	events := make([]model.Event, 0)
	err := r.storage.db.Select(
		&events,
		`SELECT title,
				       event_time,
				       duration,
				       description,
				       user_id,
				       time_to_notify
				FROM events
				WHERE event_time >= $1 and event_time <= event_time + make_interval(months => 1)`,
		date,
	)
	if err != nil {
		return nil, err
	}

	return events, nil
}
