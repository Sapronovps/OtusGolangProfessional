package model

import "time"

type Event struct {
	ID           int           `db:"id"`
	Title        string        `db:"title"`
	EventTime    time.Time     `db:"event_time"`
	Duration     time.Duration `db:"duration"`
	Description  string        `db:"description"`
	UserID       int           `db:"user_id"`
	TimeToNotify time.Time     `db:"time_to_notify"`
}
