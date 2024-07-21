/*Controlador de la tabla salas*/
package controller

import (
	"cine/DB"
	"cine/models"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"net/http"
)

// GetSalasHandler Obtiene todas las salas en la base de datos
//
//	@Summary		Obtiene salas
//	@Description	Obtencion de todas las salas en la base de datos
//	@Tags			Salas
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/api/v1/salas [get]
//
// Obtener todas las peliculas con Peliculas
func GetSalasHandler(writer http.ResponseWriter, request *http.Request) {
	var salas []models.Salas

	//encontrar y retornar todos los datos y retornoar
	//DB.DB.Find(&salas)
	DB.DB.Preload("Peliculas").Preload("Precios").Find(&salas)
	json.NewEncoder(writer).Encode(&salas)
}

// GetSalaByIDHandler Obtiene una sala en especifico
//
//	@Summary		Obtiene una sala en especifico
//	@Description	Hace la petición para obtener una sola por su ID
//
// @ID Buscar sala por su ID
// @Param ID path string true "ID de la sala a buscar"
//
//	@Tags			Salas
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/api/v1/sala/{ID} [get]
func GetSalaByIDHandler(writer http.ResponseWriter, request *http.Request) {

	//Se le dice que modelo debe de copiar para su guardado
	var sala models.Salas
	//Parametro usado para saber que se busca
	params := mux.Vars(request)

	//Obtiene los datos
	DB.DB.First(&sala, params["ID"])
	//retorna los datos
	json.NewEncoder(writer).Encode(&sala)

}
