package main

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// Тестирование
func main() {
	db, err := sqlx.Connect("postgres", "user=user password=password dbname=otus sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	//event := &Event{
	//	ID:           1,
	//	Title:        "Hello",
	//	EventTime:    time.Now(),
	//	Duration:     11,
	//	Description:  "OTUS OTUS",
	//	UserID:       1,
	//	TimeToNotify: time.Now(),
	//}
	//
	//err = CreateEvent(event, db)
	//if err != nil {
	//	log.Fatalf("Failed to create event: %v", err)
	//}
	//fmt.Println("Все ок")

	//event, err := GetEventById(1, db)
	//if err != nil {
	//	log.Fatalf("Failed to get event: %v", err)
	//}
	//event.Title = "Изменено"
	//
	//err = UpdateEvent(event, db)
	//if err != nil {
	//	log.Fatalf("Failed to update event: %v", err)
	//}

	//events, err := GetAllEvents(db)
	//
	//for event, ev := range events {
	//	_ = event
	//	fmt.Println(ev)
	//}

	err = DeleteEventById(5, db)
	if err != nil {
		log.Fatalf("Failed to delete event: %v", err)
	}
}

type Event struct {
	ID           int           `db:"id"`
	Title        string        `db:"title"`
	EventTime    time.Time     `db:"event_time"`
	Duration     time.Duration `db:"duration"`
	Description  string        `db:"description"`
	UserID       int           `db:"user_id"`
	TimeToNotify time.Time     `db:"time_to_notify"`
}

// Создание события
func CreateEvent(event *Event, db *sqlx.DB) error {
	query := `
INSERT INTO events (title, event_time, duration, description, user_id, time_to_notify)
VALUES (:title, :event_time, :duration, :description, :user_id, :time_to_notify)
`

	if _, err := db.NamedExecContext(context.Background(), query, event); err != nil {
		return err
	}
	return nil
}

// Обновление события
func UpdateEvent(event *Event, db *sqlx.DB) error {
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
	if _, err := db.NamedExecContext(context.Background(), query, event); err != nil {
		return err
	}
	return nil
}

// Получение события по ID
func GetEventById(id int, db *sqlx.DB) (*Event, error) {
	event := &Event{}
	err := db.GetContext(context.Background(), event, "SELECT * FROM events WHERE id=$1", id)
	return event, err
}

// Возвращает все события.
func GetAllEvents(db *sqlx.DB) ([]*Event, error) {
	var events []*Event
	err := db.Select(&events, "SELECT * FROM events ORDER BY event_time")
	return events, err
}

// Удаляет событие по ID.
func DeleteEventById(id int, db *sqlx.DB) error {
	_, err := db.Exec("DELETE FROM events WHERE id=$1", id)
	return err
}
