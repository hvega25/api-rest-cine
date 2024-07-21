package models

import "gorm.io/gorm"

/*Creacion de la tabla de promociones*/
type Promociones struct {
	gorm.Model
	ID          string `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Titulo      string `gorm:"not null" json:"titulo"`
	Descripcion string `gorm:"not null" json:"descripcion"`
	Url         string `gorm:"not null" json:"url"`
}
