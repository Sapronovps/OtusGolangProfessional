package memorystorage

import (
	"fmt"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/model"
	"sync"
	"time"
)

type EventRepository struct {
	events map[int]*model.Event
	mu     sync.RWMutex
}

func (r *EventRepository) Create(e *model.Event) error {
	e.ID = len(r.events) + 1
	r.mu.Lock()
	r.events[e.ID] = e
	r.mu.Unlock()
	return nil
}

func (r *EventRepository) Update(id int, e *model.Event) error {
	e.ID = id
	r.mu.Lock()
	r.events[id] = e
	r.mu.Unlock()
	return nil
}

func (r *EventRepository) Delete(id int) error {
	r.mu.Lock()
	delete(r.events, id)
	r.mu.Unlock()
	return nil
}

func (r *EventRepository) Get(id int) (model.Event, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	e, ok := r.events[id]
	if !ok {
		return model.Event{}, fmt.Errorf("error read")
	}
	return *e, nil
}

func (r *EventRepository) ListByDay(date time.Time) ([]model.Event, error) {
	events := make([]model.Event, 0)
	for _, event := range r.events {
		if event.EventTime.Format("2006-01-02") == date.Format("2006-01-02") {
			events = append(events, *event)
		}
	}
	return events, nil
}

func (r *EventRepository) ListByWeek(date time.Time) ([]model.Event, error) {
	events := make([]model.Event, 0)
	for _, event := range r.events {
		dateWeekLater := date.AddDate(0, 0, 7)
		if (event.EventTime.After(date) || event.EventTime.Equal(date)) &&
			(event.EventTime.Before(dateWeekLater) || event.EventTime.Equal(dateWeekLater)) {
			events = append(events, *event)
		}
	}
	return events, nil
}

func (r *EventRepository) ListByMonth(date time.Time) ([]model.Event, error) {
	events := make([]model.Event, 0)
	for _, event := range r.events {
		dateWeekLater := date.AddDate(0, 1, 0)
		if (event.EventTime.After(date) || event.EventTime.Equal(date)) &&
			(event.EventTime.Before(dateWeekLater) || event.EventTime.Equal(dateWeekLater)) {
			events = append(events, *event)
		}
	}
	return events, nil
}
