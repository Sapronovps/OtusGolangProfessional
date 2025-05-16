package storage

import (
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/model"
	"time"
)

type EventRepository interface {
	Create(e *model.Event) error                       // добавление события в хранилище
	Update(id int, e *model.Event) error               // изменение события в хранилище
	Delete(id int) error                               // удаление события
	Get(id int) (model.Event, error)                   // получение события
	ListByDay(date time.Time) ([]model.Event, error)   // листинг событий за день
	ListByWeek(date time.Time) ([]model.Event, error)  // листинг событий за неделю
	ListByMonth(date time.Time) ([]model.Event, error) // листинг событий за месяц
}
