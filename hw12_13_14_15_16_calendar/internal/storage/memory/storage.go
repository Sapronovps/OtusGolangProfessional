package memorystorage

import (
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/model"
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	eventRepository *EventRepository
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Event() storage.EventRepository {
	if s.eventRepository != nil {
		return s.eventRepository
	}
	s.eventRepository = &EventRepository{
		events: make(map[int]*model.Event),
	}
	return s.eventRepository
}
