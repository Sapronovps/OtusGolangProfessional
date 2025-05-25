package storage

type Storage interface {
	Event() EventRepository
}
