package sqlstorage

import (
	"github.com/Sapronovps/OtusGolangProfessional/hw12_13_14_15_calendar/internal/storage"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db              *sqlx.DB
	eventRepository *storage.EventRepository
}

func New(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) Event() storage.EventRepository {
	if s.eventRepository != nil {
		return *s.eventRepository
	}
	return &EventRepository{storage: s}
}
