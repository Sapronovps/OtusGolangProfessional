package model

import "time"

type Event struct {
	ID           int           `db:"id" json:"id"`
	Title        string        `db:"title" json:"title"`
	EventTime    time.Time     `db:"event_time" json:"-"`
	Duration     time.Duration `db:"duration" json:"-"`
	Description  string        `db:"description" json:"description"`
	UserID       int           `db:"user_id" json:"-"`
	TimeToNotify time.Time     `db:"time_to_notify" json:"-"`
}
