package controller

/*Controlador para horarios donde están definidos
todas los horarios disponibles de cada película
para obtención, edición y eliminación en la base de datos*/
import (
	"cine/DB"
	"cine/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// GetHorariosHandler Obtiene todos los horarios
//
//	@Summary		Obtiene todos los horarios disponibles
//	@Description	Obtener todos los horarios disponibles
//	@Tags			Horarios
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/api/v1/horarios [get]
func GetHorariosHandler(writer http.ResponseWriter, request *http.Request) {
	var horario []models.Horarios
	//encontrar y retornar todos los datos y retornar
	DB.DB.Find(&horario)
	//arreglar después esto
	json.NewEncoder(writer).Encode(&horario)
}

// GetHorarioByIDHandler Obtiene un horario en especifico
//
//	@Summary		Obtiene un horario en especifico
//	@Description	Hace la petición para obtener un horario por su id
//
// @id Buscar un horario por su id
// @Param id path string true "id del horario a buscar"
//
//	@Tags			Horarios
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/api/v1/horario/{id} [get]
func GetHorarioByIDHandler(writer http.ResponseWriter, request *http.Request) {

	//Se le dice que modelo debe de copiar para su guardado
	var horario []models.Horarios
	//Parametro usado para saber que se busca
	params := mux.Vars(request)

	//Obtiene los datos
	DB.DB.First(&horario, params["id"])
	//retorna los datos
	json.NewEncoder(writer).Encode(&horario)

}

// GetHorarioBySalaIDHandler Obtiene un horario por la sala donde se encuentra
//
//	@Summary		Obtiene un item por PeliculaID
//	@Description	Hace la petición para obtener un horario mediante la salaID
//
// @pelicula_id Busca llave foranea de sala en horarios
// @Param pelicula_id query string true "sala a buscar"
//
//	@Tags			Horarios
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/api/v1/horario [get]
func GetHorarioBySalaIDHandler(writer http.ResponseWriter, request *http.Request) {

	var horario []models.Horarios

	queryValues := request.URL.Query()

	//importante este user_p hace referencia de donde obtiene el parametro
	peliculaID := queryValues.Get("pelicula_id")

	//IMPORTANTE EL user_P hace referencia al model creado
	DB.DB.Where("pelicula_id = ?", peliculaID).Find(&horario)

	json.NewEncoder(writer).Encode(&horario)

}
