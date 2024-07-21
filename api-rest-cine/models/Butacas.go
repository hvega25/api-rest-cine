package models

import "gorm.io/gorm"

type Butacas struct {
	gorm.Model
	ID     uint
	Fila   string `gorm:"not null" json:"fila"`
	Numero string `gorm:"not null" json:"numero"`

	SalaID uint //llave foranea a la tabla salas
}
