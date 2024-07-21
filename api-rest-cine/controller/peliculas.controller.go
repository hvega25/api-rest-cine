/*Controlador para la rutas de peliculas */
package controller

import (
	"cine/DB"
	"cine/models"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"net/http"
)

func reversePeliculas(peliculas []models.Peliculas) {
	for i := len(peliculas)/2 - 1; i >= 0; i-- {
		opp := len(peliculas) - 1 - i
		peliculas[i], peliculas[opp] = peliculas[opp], peliculas[i]
	}
}

// GetPeliculasHandler Obtiene todas las peliculas en la base de datos
//
//	@Summary		Obtiene peliculas
//	@Description	Obtencion de todas las peliculas en la base de datos
//	@Tags			Peliculas
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/peliculas [get]
//
// Obtener todas las peliculas con Peliculas
func GetPeliculasHandler(writer http.ResponseWriter, request *http.Request) {
	var peliculas []models.Peliculas

	//encontrar y retornar todos los datos y retornoar
	DB.DB.Preload("Horarios").Find(&peliculas)
	reversePeliculas(peliculas)
	json.NewEncoder(writer).Encode(&peliculas)
}

// GetPeliculaByIDHandler Obtiene una pelicula por su ID
//
//	@Summary		Obtiene peliculas
//	@Description	Hace la petición para obtener una sola pelicula por su ID
//
// @ID Buscar elemento por su ID
// @Param id path string true "ID de la película a buscar"
//
//	@Tags			Peliculas
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/peliculas/{id} [get]
//
// Obtener una pelicula por ID
func GetPeliculaByIDHandler(writer http.ResponseWriter, request *http.Request) {

	//Se le dice a la api en donde consultar
	var pelicula models.Peliculas
	//Parametro usado para saber que se busca
	params := mux.Vars(request)

	//Obtiene los datos
	//Preloading (Eager Loading) es para cargar al json los datos asociados en la tabla
	DB.DB.Preload("Horarios").First(&pelicula, params["id"])
	//retorna los datos
	json.NewEncoder(writer).Encode(&pelicula)

}

// GetPeliculaByEstrenoHandler Obtiene una lista de películas que no sea un estreno
//
//	@Summary		Obtiene una lista de películas en cartelera
//	@Description	Hace la petición para obtener la lista de películas que estén en cartelera
//
// @estreno Encuentra las películas que en esta columna sea false
// @Param estreno query bool false "columna a comprobar"
//
//	@Tags			Peliculas
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/pelicula [get]
func GetPeliculaByEstrenoHandler(writer http.ResponseWriter, request *http.Request) {

	var pelicula []models.Peliculas

	queryValues := request.URL.Query()

	//importante este user_p hace referencia de donde obtiene el parametro
	premier := queryValues.Get("estreno")

	//IMPORTANTE EL user_P hace referencia al model creado
	DB.DB.Preload("Horarios").Where("estreno = ?", premier).Find(&pelicula)
	reversePeliculas(pelicula)
	json.NewEncoder(writer).Encode(&pelicula)

}

// CreatePeliculaHandler Creación de una película
//
//	@Summary		Insertar una película en la base de datos
//	@Description	Insertar un nueva película en la base de datos
//	@Tags			Peliculas
//	@Accept			json
//	@Produce		json
//
// @Success 201 {string} string "Éxito"
// @Failure 400 {string} string "Error"
//
//	@Router			/createpelicula [post]
func CreatePeliculaHandler(writer http.ResponseWriter, request *http.Request) {

	/*Ejemplo del json de pruebas para ingresar
		  {

	    "titulo":"Pelicula ejemplo",
	    "descripcion":"Pelicula ejemplo",
	    "sipnosisL": "Pelicula ejemplo",
	    "url": "https://ih1.redbubble.net/image.2132876762.7149/st,small,845x845-pad,1000x1000,f8f8f8.jpg",
	    "estreno": false,
	    "SaLaID": 1
	  }
	*/
	var pelicula models.Peliculas

	err := json.NewDecoder(request.Body).Decode(&pelicula)
	if err != nil {
		// Si hay un error al decodificar el cuerpo de la solicitud, responder con un código de estado 400
		http.Error(writer, "Error en los datos de la película", http.StatusBadRequest)
		return
	}
	// Intentar crear el usuario en la base de datos
	if err := DB.DB.Create(&pelicula).Error; err != nil {
		// Si hay un error al crear el usuario en la base de datos, responder con un código de estado 500 y el mensaje de error
		http.Error(writer, "Error al crear la pelicula en la base de datos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("Película creado exitosamente"))
}

// prueba
// Función para  obtener todas las películas
func GetMoviesHandler(writer http.ResponseWriter, request *http.Request) {
	/*Crea un arreglo copian la estructura que tiene el models de Movies*/
	var movies []models.Movies

	//encontrar y retornar todos los datos y retornoar
	DB.DB.Find(&movies)

	json.NewEncoder(writer).Encode(&movies)
}
