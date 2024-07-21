/*Paquete DB database */
package DB

//Librerias usadas
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

/*
DB es la variable con la que se podra obtener las querys o
funciones para el manejo de peticiones
*/
var DB *gorm.DB

// dbname=gorm
/*Variable con datos para conectar a la base de datos*/
var DSN = "host=localhost user=postgres password=passwordRoot dbname=postgres port=5432"

// Esta funcion conecta a la base de datos
func DBConnection() {
	//guarda en una variable si existe algun error
	var error error
	//intenta hacer una conexion con los parametros de la variable DSN
	/*Abre la conexión a la base de datos y si devuelve un estado 200 este se asigna a DB
	Si hay algun error este se asigna a error*/
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	/*Si hay un error lo muestra por consola y si hay conexion exitosa
	lo muestra por consola*/
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Conexión a la base de datos exitosa")
	}
}
