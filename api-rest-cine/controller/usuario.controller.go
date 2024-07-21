package controller

/*Controlador para los usuarios que se registren para poder recuperar compras,
obtención, edición y eliminación en la base de datos*/
import (
	"cine/DB"
	"cine/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// GetUsuariosHandler Usuarios de compras
//
//	@Summary		Obtención de usuarios
//	@Description	Obtención de usuarios que ya están registrados en la base de datos
//	@Tags			Usuarios
//	@Accept			json
//	@Produce		json
//
// @Success 201 {string} string "Retorno exitoso"
// @Failure 400 {string} string "Error en la obtencion de los datos"
//
//	@Router			/api/v1/usuarios [get]
func GetUsuariosHandler(writer http.ResponseWriter, request *http.Request) {
	var user []models.Usuario

	//encontrar y retornar todos los datos y retornar
	DB.DB.Preload("Tickets").Find(&user)

	//arreglar después esto
	json.NewEncoder(writer).Encode(&user)
}

// GetUsuarioByIDHandler Obtiene un usuario por su id
//
//	@Summary		Obtiene un usuario por su correo
//	@Description	Hace la petición para obtener un solo usuario por su ID
//
// @ID Busca un usuario por su ID
// @Param id path string true "ID de la usuario a buscar"
//
//	@Tags			Usuarios
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/api/v1/usuario/{id} [get]
//
// Obtener un usuario por ID
func GetUsuarioByIDHandler(writer http.ResponseWriter, request *http.Request) {

	//Se le dice a la api en donde consultar
	var usuario models.Usuario
	//Parametro usado para saber que se busca
	params := mux.Vars(request)

	//Obtiene los datos
	DB.DB.First(&usuario, params["id"])
	//retorna los datos
	json.NewEncoder(writer).Encode(&usuario)

}

// GetUsuarioByUserHandler Obtiene un usuario por su correo
//
//	@Summary		Obtiene un item por su usuario
//	@Description	Hace la petición para obtener un solo usuario por su username
//
// @user_p Busca un usuario por su usuario
// @Param user_p query string true "username del usuario a buscar"
//
//	@Tags			Usuarios
//	@Accept			json
//	@Produce		json
//	@Success		200				{string}	string
//	@Router			/api/v1/usuario [get]
//
// Falta arreglar el swagger
func GetUsuarioByUserHandler(writer http.ResponseWriter, request *http.Request) {

	var usuario models.Usuario
	/*Ejemplo query en el postman para su busqueda
	localhost:4444/api/v1/usuario?user_p=diegogomez
	usar params key user_P y diegogomez
	*/
	//query para nombre arreglar
	//DB.DB.Where("nombre = ?", "Juan").First(&usuario)
	// encuentra el usuario por correos DB.DB.Where("correo = ?", "maria@example.com").First(&usuario)
	// Si el usuario se encuentra, responde con los datos del usuario

	queryValues := request.URL.Query()

	//importante este user_p hace referencia de donde obtiene el parametro
	userP := queryValues.Get("user_p")

	//IMPORTANTE EL user_P hace referencia al model creado
	DB.DB.Where("user_P = ?", userP).First(&usuario)

	json.NewEncoder(writer).Encode(&usuario)

}

// CreateUsuariosHandler Creación de un usuario
//
//	@Summary		Insertar un usuario en la base de datos
//	@Description	Insertar un nuevo usuario en la base de datos
//	@Tags			Usuarios
//	@Accept			json
//	@Produce		json
//
// @Success 201 {string} string "Retorno exitoso"
// @Failure 400 {string} string "Error en la obtencion de los datos"
//
//	@Router			/api/v1/usuario [post]
func CreateUsuariosHandler(writer http.ResponseWriter, request *http.Request) {

	/*Ejemplo del json de pruebas para ingresar
		{
	    "nombre": "Herson",
	    "apellido": "Vega",
	    "user_P": "hvega25",
	    "correo": "hvega@example.com",
	    "contrasena": "contraseña123",
	    "nacimiento": "1987-04-12T00:00:00Z",
	    "sexo": "mosca"
	}
	*/
	var user models.Usuario

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		// Si hay un error al decodificar el cuerpo de la solicitud, responder con un código de estado 400
		http.Error(writer, "Error en los datos del usuario", http.StatusBadRequest)
		return
	}
	// Intentar crear el usuario en la base de datos
	if err := DB.DB.Create(&user).Error; err != nil {
		// Si hay un error al crear el usuario en la base de datos, responder con un código de estado 500 y el mensaje de error
		http.Error(writer, "Error al crear el usuario en la base de datos: "+err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte("Usuario creado exitosamente"))
}
