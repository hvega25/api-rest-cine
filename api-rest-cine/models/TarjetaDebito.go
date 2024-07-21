package models

import "gorm.io/gorm"

/*Tabla de las tarjetas de debitos los cuales en ella se van a consultar
 */
type TarjetaDebito struct {
	gorm.Model

	Titular        string `gorm:"not null" json:"titular"`
	Numero_Tarjeta string `gorm:"not null" json:"numero___tarjeta"`
	CVV            string `gorm:"not null" json:"cvv"`
	Expira         string `gorm:"not null" json:"expira"`
	Saldo          int64  `gorm:"not null" json:"saldo"`
}
