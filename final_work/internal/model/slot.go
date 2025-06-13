package model

import (
	"sync"
	"time"
)

type Slot struct {
	mu          sync.Mutex
	ID          int                 `db:"id"`          // ID слота
	Description string              `db:"description"` // Описание слота
	Banners     map[int]*Banner     // Все баннеры слота
	DecayLambda float64             `db:"decay_lambda"` // Коэффициент затухания
	HistorySize int                 `db:"history_size"` // Размер окна для учета усталости
	ShowHistory map[int][]time.Time // История показов для учета усталости
}
