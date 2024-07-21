package models

import "gorm.io/gorm"

type Salas struct {
	gorm.Model

	Nombre      string `gorm:"not null" json:"nombre"`
	Descripcion string `gorm:"not null" json:"descripcion"`
	Capacidad   string `gorm:"not null" json:"capacidad"`
	Url         string `gorm:"not null" json:"url"`

	//Peliculas
	Peliculas []Peliculas `gorm:"foreignKey:SalaID"`

	//Precios
	Precios []Precios `gorm:"foreignKey:SalaID"`
}
