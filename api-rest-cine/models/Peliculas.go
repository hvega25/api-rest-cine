package models

import (
	"gorm.io/gorm"
)

/*Estructura de la tabla*/
type Peliculas struct {
	gorm.Model

	Titulo      string `gorm:"not null" json:"titulo"`
	Descripcion string `gorm:"not null" json:"descripcion"`
	SipnosisL   string `gorm:"not null" json:"sipnosisL"`
	Url         string `gorm:"not null" json:"url"`

	//Prueba con columna nueva para determinar si es estreno o no
	Estreno bool `json:"estreno"`
	/*Nota el foreignKey:PeliculaID se muestra en horarios */
	Horarios []Horarios `gorm:"foreignKey:PeliculaID"` // Relación uno a muchos con los horarios

	//Union con salas
	SalaID uint
}
