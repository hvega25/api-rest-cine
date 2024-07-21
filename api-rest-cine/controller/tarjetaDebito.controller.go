package controller

/*Controlador para las tarjetas de débito donde se obtienen
unas tarjetas como simulador para poder hacer una compra,
obtención, edición y eliminación en la base de datos*/
import (
	"cine/DB"
	"cine/models"
	"encoding/json"
	"net/http"
)

// GetTarjetaDebitoHandler Obtiene todas las promociones
//
//	@Summary		Obtener las tarjetas de débito
//	@Description	Obtener tarjetas de débito o crédito para simular una compra
//	@Tags			Tarjetas
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/api/v1/TarjetasDebito [get]
func GetTarjetaDebitoHandler(writer http.ResponseWriter, request *http.Request) {
	var tarjetas []models.TarjetaDebito

	//encontrar y retornar todos los datos y retornar
	DB.DB.Find(&tarjetas)

	//arreglar después esto
	json.NewEncoder(writer).Encode(&tarjetas)
}

// GetTarjeaByNumberHandler Obtiene un usuario por su correo
//
//	@Summary		Obtiene un item por su numero de tarjeta
//	@Description	Hace la petición para obtener un sola tarjeta por numero
//
// @cardNumber Busca una tarjeta por su numero
// @Param cardNumber query string true "Numero de la tarjeta a buscar"
//
//	@Tags			Tarjetas
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/api/v1/tarjeta [get]
func GetTarjeaByNumberHandler(writer http.ResponseWriter, request *http.Request) {

	var tarjeta models.TarjetaDebito
	/*Ejemplo query en el postman para su busqueda
	localhost:4444/api/v1/usuario?user_p=diegogomez
	usar params key user_P y diegogomez
	*/
	//query para nombre arreglar
	//DB.DB.Where("nombre = ?", "Juan").First(&usuario)
	// encuentra el usuario por correos DB.DB.Where("correo = ?", "maria@example.com").First(&usuario)
	// Si el usuario se encuentra, responde con los datos del usuario

	queryValues := request.URL.Query()
	debit := queryValues.Get("cardNumber")
	DB.DB.Where("Numero_Tarjeta = ?", debit).First(&tarjeta)

	json.NewEncoder(writer).Encode(&tarjeta)

}
