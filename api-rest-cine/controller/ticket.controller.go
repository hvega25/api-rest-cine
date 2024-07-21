package controller

import (
	"cine/DB"
	"cine/models"
	"encoding/json"
	"net/http"
)

// GetTicketHandler Usuarios de compras
//
//	@Summary		Obtiene todas las tickets
//	@Description	Obtiene todos los tickets ya sea de un usuario registrado o no
//	@Tags			Tickets
//	@Accept			json
//	@Produce		json
//
// @Success 200 {string} string "Exitoso"
// @Failure 400 {string} string "Error en la obtención de tickets"
//
//	@Router			/api/v1/tickets [get]
func GetTicketHandler(w http.ResponseWriter, r *http.Request) {
	var ticket []models.Ticket

	DB.DB.Find(&ticket)

	json.NewEncoder(w).Encode(&ticket)

}

// CreateTicketHandler Creación de un TICKET
//
//	@Summary		Insertar un ticket en la base de datos
//	@Description	Insertar un nuevo ticket en la base de datos
//	@Tags			Tickets
//	@Accept			json
//	@Produce		json
//
// @Success 201 {string} string "Creación exitosa"
// @Failure 400 {string} string "Error en el retorno"
//
//	@Router			/api/v1/ticket [post]
func CreateTicketHandler(writer http.ResponseWriter, request *http.Request) {
	var ticket models.Ticket

	/*EJEMPLO TICKET JSON
		{
	    "codigo": "Herson",
	    "butaca": "Vega",
	    "precio": "hvega25",
	    "hora": "hvega@example.com",
	    "idpelicula": 2,
	    "tituloPelicula": "1987-04-12T00:00:00Z",
	    "usuario_id": 1
	}
	*/
	err := json.NewDecoder(request.Body).Decode(&ticket)
	if err != nil {
		// Si hay un error al decodificar el cuerpo de la solicitud, responder con un código de estado 400
		http.Error(writer, "Error en los datos del ticket", http.StatusBadRequest)
		return
	}
	// Intentar crear el usuario en la base de datos
	if err := DB.DB.Create(&ticket).Error; err != nil {
		// Si hay un error al crear el usuario en la base de datos, responder con un código de estado 500 y el mensaje de error
		http.Error(writer, "Error al crear el ticket en la base de datos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("Ticket creado exitosamente"))
}
