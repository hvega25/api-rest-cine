package models

/*El uso de la librería gorm facilita la extracción, manejo de datos
y migración hacia la base de datos*/
import (
	"gorm.io/gorm"
)

/*Estructura de la tabla */
/*NOTA PARA EL FUTURO PARA CREAR LA TABLA SIEMPRE LA PRIMERA LETRA USAR MAYUSCULA*/
type Movies struct {
	gorm.Model

	/*En la primera columna se declara la variable que se migrara IMPORTANTE PRIMERA EN MAYÚSCULA*/
	/*Segunda columna indica de que tipo se va a manejar el valor*/
	/*En la tercera columna hace uso del gorm y le indica que se debe de comportar*/
	Titulo      string `gorm:"not null"`
	Descripcion string `gorm:"not null"`
	SipnosisL   string `gorm:"not null"`
	Url         string `gorm:"not null"`
}
