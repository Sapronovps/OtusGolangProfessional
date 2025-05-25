package memorystorage

import (
	"fmt"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/model"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCreateAndGet(t *testing.T) {
	storage, testEvents := fillStorage()

	for i, event := range testEvents {
		i++
		event.ID = i
		createdEvent, err := storage.Event().Get(i)
		require.Nil(t, err)
		require.Equal(t, event, createdEvent)
	}
}

func TestUpdate(t *testing.T) {
	storage, _ := fillStorage()
	eventID := 2
	event, err := storage.Event().Get(eventID)
	require.Nil(t, err)
	require.Equal(t, event.Title, "Event 2")

	newTitle := "NewEvent 2"
	event.Title = newTitle
	err = storage.Event().Update(eventID, &event)
	require.Nil(t, err)

	event, err = storage.Event().Get(eventID)
	require.Nil(t, err)
	require.Equal(t, event.Title, newTitle)
}

func TestDelete(t *testing.T) {
	storage, _ := fillStorage()
	eventID := 2
	event, err := storage.Event().Get(eventID)
	require.Nil(t, err)
	require.NotNil(t, event)

	err = storage.Event().Delete(eventID)
	require.Nil(t, err)
	_, err = storage.Event().Get(eventID)
	require.NotNil(t, err)
}

func getTestEvents() []model.Event {
	events := make([]model.Event, 5)

	for i := 0; i < 5; i++ {
		events[i] = model.Event{
			ID:        i + 1,
			Title:     fmt.Sprintf("Event %d", i+1),
			EventTime: time.Now().AddDate(0, 0, 7),
			UserID:    i,
		}
	}

	return events
}

func fillStorage() (*Storage, []model.Event) {
	storage := New()
	testEvents := getTestEvents()
	for _, event := range testEvents {
		_ = storage.Event().Create(&event)
	}

	return storage, testEvents
}
