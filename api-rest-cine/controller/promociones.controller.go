package controller

/*Controlador para promociones donde están definidos
todos los promociones disponibles
para obtención, edición y eliminación en la base de datos*/
import (
	"cine/DB"
	"cine/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// GetPromocionesHandler Obtiene todas las promociones
//
//	@Summary		Obtener todos los promociones
//	@Description	Obtener todas las promociones disponibles
//	@Tags			Promociones
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/api/v1/promociones [get]
func GetPromocionesHandler(writer http.ResponseWriter, request *http.Request) {
	var promociones []models.Promociones

	//encontrar y retornar todos los datos y retornar
	DB.DB.Find(&promociones)

	//arreglar después esto
	json.NewEncoder(writer).Encode(&promociones)
}

// GetPromocionByIDHandler Obtiene una promocion en especifico
//
//	@Summary		Obtiene una promocion en especifico
//	@Description	Hace la petición para obtener una promocion por su id
//
// @id Buscar una promocion por su id
// @Param id path string true "ID de la promocion a buscar"
//
//	@Tags			Promociones
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/api/v1/promocion/{id} [get]
func GetPromocionByIDHandler(writer http.ResponseWriter, request *http.Request) {

	//Se le dice que modelo debe de copiar para su guardado
	var promo models.Promociones
	//Parametro usado para saber que se busca
	params := mux.Vars(request)

	//Obtiene los datos
	DB.DB.First(&promo, params["id"])
	//retorna los datos
	json.NewEncoder(writer).Encode(&promo)

}
