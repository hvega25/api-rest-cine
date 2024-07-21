package models

import "gorm.io/gorm"

type Precios struct {
	gorm.Model

	//Llave foranea de ticket
	//TicketID uint `gorm:"unique; not null" json:"ticket_id"`
	//Lllave foranea de Salas
	SalaID uint ` json:"sala_id"`

	//Columnas propias de Precios
	Precio float64 `gorm:"not null" json:"precio"`
	Moneda string  `gorm:"not null" json:"moneda"`
}
