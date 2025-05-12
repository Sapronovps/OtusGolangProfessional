package app

import (
	"context"
)

type App struct {
	Logger
	Storage
}

type Logger interface {
	Error(msg string)
	Warning(msg string)
	Info(msg string)
	Debug(msg string)
}

type Storage interface {
	Create(ctx context.Context) error
}

func New(logger Logger, storage Storage) *App {
	return &App{}
}

func (a *App) CreateEvent(ctx context.Context, id, title string) error {
	// TODO
	return nil
	// return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

// TODO
