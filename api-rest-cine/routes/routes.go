package routes

/*Archivo para el guardado de rutas, este contiene todas las rutas
de los modelos que se quiere se escuchen en el localhost*/
import (
	"cine/controller"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func Rutas(router *mux.Router) {

	//Ruta prueba para saber si funciona
	router.HandleFunc("/", controller.Home)

	//Rutas para peliculas
	router.HandleFunc("/peliculas", controller.GetPeliculasHandler).Methods("GET").Name("peliculas")
	router.HandleFunc("/createpelicula", controller.CreatePeliculaHandler).Methods("POST").Name("insertpelicula")
	router.HandleFunc("/pelicula", controller.GetPeliculaByEstrenoHandler).Methods("GET").Name("peliculasEstrenoFalse")
	router.HandleFunc("/movies", controller.GetMoviesHandler).Methods("GET").Name("movies")
	router.HandleFunc("/peliculas/{id}", controller.GetPeliculaByIDHandler).Methods("GET").Name("peliculaID")

	//Rutas para salas
	router.HandleFunc("/api/v1/salas", controller.GetSalasHandler).Methods("GET").Name("salas")
	router.HandleFunc("/api/v1/sala/{ID}", controller.GetSalaByIDHandler).Methods("GET").Name("sala")

	//Rutas para butacas
	router.HandleFunc("/api/v1/butacas", controller.GetButacasHandler).Methods("GET").Name("butacas")
	router.HandleFunc("/api/v1/butaca", controller.GetButacaBySalaIDHandler).Methods("GET").Name("butacasSalaID")
	//idea buscar por fila y numero

	//Rutas para horarios
	router.HandleFunc("/api/v1/horarios", controller.GetHorariosHandler).Methods("GET").Name("horarios")
	router.HandleFunc("/api/v1/horario/{id}", controller.GetHorarioByIDHandler).Methods("GET").Name("horario")
	router.HandleFunc("/api/v1/horario", controller.GetHorarioBySalaIDHandler).Methods("GET").Name("horariosPeliculaID")

	//Rutas para precios
	router.HandleFunc("/api/v1/precios", controller.GetPreciosHandler).Methods("GET").Name("precios")
	router.HandleFunc("/api/v1/precio/{id}", controller.GetPrecioByIDHandler).Methods("GET").Name("precio")
	router.HandleFunc("/api/v1/precio", controller.GetPrecioBySalaIDHandler).Methods("GET").Name("precioSalaID")

	//Rutas para promociones
	router.HandleFunc("/api/v1/promociones", controller.GetPromocionesHandler).Methods("GET").Name("promociones")
	router.HandleFunc("/api/v1/promocion/{id}", controller.GetPromocionByIDHandler).Methods("GET").Name("promocion")

	//Rutas para promociones
	router.HandleFunc("/api/v1/TarjetasDebito", controller.GetTarjetaDebitoHandler).Methods("GET").Name("tarjetas")
	router.HandleFunc("/api/v1/tarjeta", controller.GetTarjeaByNumberHandler).Methods("GET").Name("numbercard")

	//Rutas para usuarios
	router.HandleFunc("/api/v1/usuarios", controller.GetUsuariosHandler).Methods("GET").Name("usuarios")
	router.HandleFunc("/api/v1/usuario/{id}", controller.GetUsuarioByIDHandler).Methods("GET").Name("usuario")
	router.HandleFunc("/api/v1/usuario", controller.GetUsuarioByUserHandler).Methods("GET").Name("User")
	router.HandleFunc("/api/v1/usuario", controller.CreateUsuariosHandler).Methods("POST").Name("User")

	//Rutas para tickets
	router.HandleFunc("/api/v1/tickets", controller.GetTicketHandler).Methods("GET").Name("tickets")
	router.HandleFunc("/api/v1/ticket", controller.CreateTicketHandler).Methods("POST").Name("ticketCreado")

	//swagger
	/*Siempre recordar que para agregar una ruta a swagger
	ejecutar este comando en consola: swag init -g main.go*/
	router.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.json")
	})
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

}
