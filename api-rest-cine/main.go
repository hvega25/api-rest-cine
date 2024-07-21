// Paquete principal IMPORTANTE PARA EJECUTAR
package main

/*Cine/DB significa que cine es el modulo creado y DB el paquete*/
import (
	"cine/DB"
	_ "cine/docs"
	"cine/models"
	"cine/routes"
	"cine/script"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

// @title			Cine
// @version		1.0
// @description	Api rest en GO de un cine
// @host			localhost:4444
// @BasePath		/
func main() {

	//ejecutando la conexión al iniciar el programa
	DB.DBConnection()
	script.EliminarTablas()

	//script.CrearTablas()
	/*Crea la variable usada para los endpoints usando Gorrilla / MUX como
	libreria para el mapeo de peticiones entrantes*/
	router := mux.NewRouter()
	/*Función que contiene todas las rutas de la API
	Esto con la intensión de modular mas el código*/
	routes.Rutas(router)

	//Integrando el cors para que sean globales en toda función que lo necesite
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Permitir solicitudes desde este origen
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(router)
	//ejecuta la migración automática de la base de datos
	//INICIAR CON EL MIGRADO DE LOS MODELS
	//el primer db hace referencia al paquete
	/*El segundo db hace referencia al gorm que este contiene un repositorio de querys para
	la base de datos*/
	DB.DB.AutoMigrate(&models.Usuario{})
	DB.DB.AutoMigrate(&models.Peliculas{})
	DB.DB.AutoMigrate(&models.Salas{})
	DB.DB.AutoMigrate(&models.Butacas{})
	DB.DB.AutoMigrate(&models.Horarios{})
	DB.DB.AutoMigrate(&models.Ticket{})
	DB.DB.AutoMigrate(&models.Precios{})
	DB.DB.AutoMigrate(&models.TarjetaDebito{})
	DB.DB.AutoMigrate(&models.Promociones{})

	script.InsertarDatos()

	/*http es la libreria net/http propia de GO que usa los protocolos de internet
	Función que inicia la API este con 2 parametros
	El primer parametro indica en que puerto esta disponible
	El segundo parametro hace referencia a la libreria Gorilla / MUX que maneja las peticiones entrantes */
	http.ListenAndServe(":4444", handler)

}

/*
Manifiesto

Programador, si has llegado hasta este código, es porque estás desarrollando una api rest o simplemente tienes una pequeña
curiosidad de saber qué hay; y si planeas usarlo, toma en cuenta los siguientes puntos:

 Primero
Sentite libre de usar este código, tienes el permiso para hacer uso.

 Segundo
Usé la librería de swaggo para mux para hacer el swagger en su momento, lo hice con este comando en el proyecto
go install github.com/swaggo/swag/cmd/swag@latest
swag init, este comando ejecutarlo cada vez que agregues una nueva ruta

También, hice uso de un liveserver que es una librería
go install github.com/cosmtrek/air@latest para instalarlo
air init, para iniciarlo
air, para ejecutarlo

Se ha hecho el uso de cors globales usando este comando
go get github.com/rs/cors

 Tercero
Cuando escribía este código, dios y yo sabíamos de que iba,
ahora solo dios lo sabe */
