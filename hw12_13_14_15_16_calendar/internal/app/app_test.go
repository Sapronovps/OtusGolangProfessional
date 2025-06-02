package app

import (
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/model"
	memorystorage "github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/storage/memory"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestCalendarCRUD(t *testing.T) {
	logg := zap.NewExample()
	storage := memorystorage.New()
	app := New(logg, storage)

	// Создадим событие.
	event := model.Event{
		Title:       "Тестовое событие",
		Description: "Проверка события",
		EventTime:   time.Now(),
	}
	err := app.CreateEvent(&event)

	require.Nil(t, err)

	// Получим событие.
	event1, err := app.GetEvent(1)

	require.Nil(t, err)
	require.Equal(t, event.Title, event1.Title)

	// Обновим событие.
	newTitle := "Обновленное событие"
	event1.Title = newTitle
	err = app.UpdateEvent(1, &event1)

	require.Nil(t, err)

	updatedEvent, err := app.GetEvent(1)

	require.Nil(t, err)
	require.Equal(t, newTitle, updatedEvent.Title)

	// Получим события за день.
	events, err := app.ListByDay(time.Now())

	require.Nil(t, err)
	require.Len(t, events, 1)
	expectedEvent := []model.Event{event1}
	require.Equal(t, expectedEvent, events)

	// Удалим событие.
	err = app.DeleteEvent(1)
	require.Nil(t, err)

	// Попробуем получить удаленное событие.
	_, err = app.GetEvent(1)
	require.NotNil(t, err)
	require.Equal(t, "error read", err.Error())
}
