package models

import (
	"gorm.io/gorm"
	"time"
)

type Usuario struct {
	gorm.Model

	Nombre     string    `gorm:"not null" json:"nombre"`
	Apellido   string    `gorm:"not null" json:"apellido"`
	User_p     string    `json:"user_P"`
	Correo     string    `gorm:"not null" json:"correo"`
	Contrasena string    `gorm:"not null" json:"contrasena"`
	Nacimiento time.Time `gorm:"not null" json:"nacimiento"`
	Sexo       string    `gorm:"not null" json:"sexo"`

	Tickets []Ticket `json:"tickets"`
}
