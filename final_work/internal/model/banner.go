package model

import "time"

type Banner struct {
	ID          int       `db:"id"`          // ID баннера
	Description string    `db:"description"` // Описание баннера
	Shows       int       `db:"shows"`       // Количество показов
	Clicks      int       `db:"clicks"`      // Количество кликов
	LastShown   time.Time `db:"last_shown"`  // Время последнего показа
	Weight      float64   `db:"weight"`      // Текущий вес (CTR или другая метрика)
}
