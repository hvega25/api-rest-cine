package controller

import "net/http"

// Home HealthCheck
//
//	@Summary		Saber si el servidor esta en ejecución
//	@Description	Prueba rápida para saber si el servidor está en ejecución
//	@Tags			HealthCheck
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/ [get]
func Home(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hola mundo!"))
}
