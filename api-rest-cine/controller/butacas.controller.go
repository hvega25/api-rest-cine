package controller

/*Controlador para butacas donde estan definidos
todas las rutas para obtención, edición y eliminación en la base de datos*/
import (
	"cine/DB"
	"cine/models"
	"encoding/json"
	"net/http"
)

// GetButacasHandler Obtiene todas las butacas de cualquier sala
//
//	@Summary		Obtiene todas las
//	@Description	Obtención de todas las butacas en la base de datos
//	@Tags			Butacas
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/api/v1/butacas [get]
func GetButacasHandler(writer http.ResponseWriter, request *http.Request) {
	var butacas []models.Butacas

	//encontrar y retornar todos los datos y retornar
	DB.DB.Find(&butacas)

	//arreglar después esto
	json.NewEncoder(writer).Encode(&butacas)
}

// GetButacaBySalaIDHandler Obtiene las butacas que se encuentran en una sala por medio de la SalaID
//
//	@Summary		Obtiene un item por salaID
//	@Description	Hace la petición para obtener las butacas mediante la salaID
//
// @sala_id Busca llave foranea de sala en butaca
// @Param sala_id query int true "clave foranea a buscar"
//
//	@Tags			Butacas
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/api/v1/butaca [get]
func GetButacaBySalaIDHandler(writer http.ResponseWriter, request *http.Request) {

	var butacas []models.Butacas

	queryValues := request.URL.Query()

	roomID := queryValues.Get("sala_id")
	if roomID == "" {
		// Manejar el caso en el que roomID está vacío
		http.Error(writer, "El parámetro sala_id es requerido", http.StatusBadRequest)
		return
	}
	//IMPORTANTE EL user_P hace referencia al model creado
	DB.DB.Where("sala_id = ?", roomID).Find(&butacas)

	json.NewEncoder(writer).Encode(&butacas)

}
