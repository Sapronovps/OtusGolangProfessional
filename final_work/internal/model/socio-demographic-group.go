package model

type SocioDemographicGroup struct {
	ID          int    `db:"id"`
	Description string `db:"description"`
}
