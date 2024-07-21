package controller

/*Controlador para precios donde están definidos
todas los precios disponibles de cada película
para obtención, edición y eliminación en la base de datos*/
import (
	"cine/DB"
	"cine/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// GetPreciosHandler Obtiene todos los precios de todas las películas
//
//	@Summary		Obtener todos los precios
//	@Description	Obtener todos los precios
//	@Tags			Precios
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/api/v1/precios [get]
func GetPreciosHandler(writer http.ResponseWriter, request *http.Request) {
	var precios []models.Precios

	//encontrar y retornar todos los datos y retornar
	DB.DB.Find(&precios)

	//arreglar después esto
	json.NewEncoder(writer).Encode(&precios)
}

// GetPrecioByIDHandler Obtiene un precio en especifico
//
//	@Summary		Obtiene un precio en especifico
//	@Description	Hace la petición para obtener una solo precio por su ID
//
// @ID Buscar precio por su ID
// @Param id path string true "ID del precio a buscar"
//
//	@Tags			Precios
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/api/v1/precio/{id} [get]
func GetPrecioByIDHandler(writer http.ResponseWriter, request *http.Request) {

	//Se le dice que modelo debe de copiar para su guardado
	var precio models.Precios
	//Parametro usado para saber que se busca
	params := mux.Vars(request)

	//Obtiene los datos
	DB.DB.First(&precio, params["id"])
	//retorna los datos
	json.NewEncoder(writer).Encode(&precio)

}

// GetPrecioBySalaIDHandler Función que retorna los precios asociados a una sala
//
//	@Summary		Obtiene precios con la llave foranea SalaID
//	@Description	Hace la petición para obtener los precios mediante la salaID
//
// @sala_id Busca llave foranea de sala en precios
// @Param sala_id query int true "llave forenea de sala en precios"
//
//	@Tags			Precios
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/api/v1/precio [get]
func GetPrecioBySalaIDHandler(writer http.ResponseWriter, request *http.Request) {

	var precio models.Precios

	queryValues := request.URL.Query()

	roomID := queryValues.Get("sala_id")
	if roomID == "" {
		// Manejar el caso en el que roomID está vacío
		http.Error(writer, "El parámetro sala_id es requerido", http.StatusBadRequest)
		return
	}
	//IMPORTANTE EL user_P hace referencia al model creado
	DB.DB.Where("sala_id = ?", roomID).Find(&precio)

	json.NewEncoder(writer).Encode(&precio)

}
