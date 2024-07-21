package models

import (
	"gorm.io/gorm"
)

/*Estructura nueva de tickets*/
type Ticket struct {
	gorm.Model

	Codigo         string `gorm:"unique" json:"codigo"`
	Butaca         string `gorm:"not null" json:"butaca"`
	Precio         string `gorm:"not null" json:"precio"`
	Hora           string `gorm:"not null" json:"hora"`
	IDpelicula     int    `gorm:"not null" json:"idpelicula"`
	TituloPelicula string `gorm:"not null" json:"tituloPelicula"`
	//llave foranea
	UsuarioID *uint ` json:"usuario_id"`
}
