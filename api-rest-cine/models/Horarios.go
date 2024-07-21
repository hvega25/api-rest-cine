package models

import "gorm.io/gorm"

type Horarios struct {
	gorm.Model

	Fecha string `gorm:"not null" json:"fecha"`
	Hora  string `gorm:"not null" json:"hora"`
	/*el PeliculaID viene de Peliculas*/
	PeliculaID uint ` json:"pelicula_id"` // Campo de referencia a la película
}
