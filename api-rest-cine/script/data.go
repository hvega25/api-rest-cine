package script

import (
	"bufio"
	"cine/DB"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
)

func EliminarTablas() {
	DB.DB.Exec("DROP TABLE IF EXISTS Peliculas CASCADE")
	fmt.Println("Peliculas eliminada")

	DB.DB.Exec("DROP TABLE IF EXISTS Butacas CASCADE")
	fmt.Println("Butacas eliminada")

	DB.DB.Exec("DROP TABLE IF EXISTS Horarios CASCADE")
	fmt.Println("Horarios eliminada")

	DB.DB.Exec("DROP TABLE IF EXISTS Precios CASCADE")
	fmt.Println("Precios eliminada")

	DB.DB.Exec("DROP TABLE IF EXISTS Promociones CASCADE")
	fmt.Println("Promociones eliminada")

	DB.DB.Exec("DROP TABLE IF EXISTS Salas CASCADE")
	fmt.Println("Salas eliminada")

	DB.DB.Exec("DROP TABLE IF EXISTS Tarjeta_debitos CASCADE")
	fmt.Println("Tarjetas eliminada")

	DB.DB.Exec("DROP TABLE IF EXISTS Tickets CASCADE")
	fmt.Println("Ticket eliminada")

	DB.DB.Exec("DROP TABLE IF EXISTS Usuarios CASCADE")
	fmt.Println("Usuarios eliminada")

	DB.DB.Exec("DROP TABLE IF EXISTS Movies CASCADE")
	fmt.Println("Tabla de prueba eliminada eliminada")

}

func CrearTablas() {

	DB.DB.Exec("CREATE TABLE Peliculas (id SERIAL PRIMARY KEY)")
	fmt.Println("Peliculas creada")

	DB.DB.Exec("CREATE TABLE Butacas (id SERIAL PRIMARY KEY)")
	fmt.Println("Butacas creada")

	DB.DB.Exec("CREATE TABLE Horarios (id SERIAL PRIMARY KEY)")
	fmt.Println("Horarios creada")

	DB.DB.Exec("CREATE TABLE Precios (id SERIAL PRIMARY KEY)")
	fmt.Println("Precios creada")

	DB.DB.Exec("CREATE TABLE Promociones (id SERIAL PRIMARY KEY)")
	fmt.Println("Promociones creada")

	DB.DB.Exec("CREATE TABLE Salas (id SERIAL PRIMARY KEY)")
	fmt.Println("Salas creada")

	DB.DB.Exec("CREATE TABLE TarjetaDebito (id SERIAL PRIMARY KEY)")
	fmt.Println("Tarjeta Debito creada")

	DB.DB.Exec("CREATE TABLE Ticket (id SERIAL PRIMARY KEY)")
	fmt.Println("ticket  creada")

	DB.DB.Exec("CREATE TABLE Usuarios (id SERIAL PRIMARY KEY)")
	fmt.Println("Usuarios creada")
}

/*Función que lee los datos desde un archivo txt*/
func InsertarDatosDesdeArchivo(nombreArchivo string) error {

	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return err
	}
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		valores := strings.Split(linea, ",")
		if len(valores) != 8 {
			return fmt.Errorf("Formato incorrecto en línea: %s", linea)
		}

		// Eliminar espacios en blanco al inicio y al final de cada valor
		for i := range valores {
			valores[i] = strings.TrimSpace(valores[i])
		}

		// Construir la consulta SQL
		query := fmt.Sprintf("INSERT INTO Salas (id, created_at, updated_at, deleted_at, nombre, descripcion, capacidad, url) VALUES (%s, %s, %s, %s, '%s', '%s', %s, '%s')",
			valores[0], valores[1], valores[2], valores[3], valores[4], valores[5], valores[6], valores[7])

		// Ejecutar la consulta SQL
		DB.DB.Exec(query)
	}

	fmt.Println("Inserciones completadas desde el archivo:", nombreArchivo)
	return nil
}

/*Función que lee los datos desde un archivo txt*/
func InsertarPeliculas(nombreArchivo string) error {

	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return err
	}
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		valores := strings.Split(linea, ",")
		if len(valores) != 8 {
			return fmt.Errorf("Formato incorrecto en línea: %s", linea)
		}

		// Eliminar espacios en blanco al inicio y al final de cada valor
		for i := range valores {
			valores[i] = strings.TrimSpace(valores[i])
		}

		// Construir la consulta SQL
		query := fmt.Sprintf("INSERT INTO Peliculas ( created_at, updated_at, titulo, descripcion, sipnosis_l, url, estreno, sala_id) VALUES (%s, %s, '%s', '%s', '%s','%s',%s ,%s)",
			valores[0], valores[1], valores[2], valores[3], valores[4], valores[5], valores[6], valores[7])

		// Ejecutar la consulta SQL
		DB.DB.Exec(query)

	}

	fmt.Println("Inserciones completadas desde el archivo:", nombreArchivo)
	return nil
}

/*Función que lee los datos desde un archivo txt*/
func InsertarUsuarios(nombreArchivo string) error {

	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return err
	}
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		valores := strings.Split(linea, ",")
		if len(valores) != 9 {
			return fmt.Errorf("Formato incorrecto en línea: %s", linea)
		}

		// Eliminar espacios en blanco al inicio y al final de cada valor
		for i := range valores {
			valores[i] = strings.TrimSpace(valores[i])
		}

		//Generando los hashes antes de guardarlo
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(valores[6]), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("Error al generar el hash de la contraseña: %v", err)
		}

		// Construir la consulta SQL
		query := fmt.Sprintf("INSERT INTO Usuarios (created_at, updated_at, nombre, apellido, user_p, correo, contrasena, nacimiento, sexo) VALUES ( %s, %s, '%s', '%s', '%s','%s','%s','%s','%s')",
			valores[0], valores[1], valores[2], valores[3], valores[4], valores[5], string(hashedPassword), valores[7], valores[8])

		// Ejecutar la consulta SQL
		DB.DB.Exec(query)

	}

	fmt.Println("Inserciones completadas desde el archivo:", nombreArchivo)
	return nil
}

/*Función que lee los datos desde un archivo txt*/
func InsertarTarjeta(nombreArchivo string) error {

	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return err
	}
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		valores := strings.Split(linea, ",")
		if len(valores) != 9 {
			return fmt.Errorf("Formato incorrecto en línea: %s", linea)
		}

		// Eliminar espacios en blanco al inicio y al final de cada valor
		for i := range valores {
			valores[i] = strings.TrimSpace(valores[i])
		}

		// Construir la consulta SQL
		query := fmt.Sprintf("INSERT INTO Tarjeta_debitos (id, created_at, updated_at, deleted_at, titular, numero_tarjeta, cvv, expira, saldo) VALUES (%s, %s, %s, %s, '%s', '%s','%s','%s',%s)",
			valores[0], valores[1], valores[2], valores[3], valores[4], valores[5], valores[6], valores[7], valores[8])

		// Ejecutar la consulta SQL
		DB.DB.Exec(query)

	}

	fmt.Println("Inserciones completadas desde el archivo:", nombreArchivo)
	return nil
}

// Insertar butaca
func InsertarButacas() error {
	salas := []int{1, 2, 3, 4}
	for _, sala := range salas {
		for i := 1; i <= 20; i++ {
			id := (sala-1)*20 + i
			query := fmt.Sprintf("INSERT INTO butacas(id, fila, numero, sala_id) VALUES (%d, 'A', %d, %d)", id, i, sala)
			DB.DB.Exec(query)
		}
	}

	fmt.Println("Inserción de butacas completada")
	return nil
}

/*Función que lee los datos desde un archivo txt*/
func InsertarHorarios(nombreArchivo string) error {

	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return err
	}
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		valores := strings.Split(linea, ",")
		if len(valores) != 6 {
			return fmt.Errorf("Formato incorrecto en línea: %s", linea)
		}

		// Eliminar espacios en blanco al inicio y al final de cada valor
		for i := range valores {
			valores[i] = strings.TrimSpace(valores[i])
		}

		// Construir la consulta SQL
		query := fmt.Sprintf("INSERT INTO horarios(created_at, updated_at, deleted_at, fecha, hora, pelicula_id) VALUES (%s, %s, %s,'%s','%s', %s)",
			valores[0], valores[1], valores[2], valores[3], valores[4], valores[5])

		// Ejecutar la consulta SQL
		DB.DB.Exec(query)

	}

	fmt.Println("Inserciones completadas desde el archivo:", nombreArchivo)
	return nil
}

/*Función que lee los datos desde un archivo txt*/
func InsertarPrecios(nombreArchivo string) error {

	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return err
	}
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		valores := strings.Split(linea, ",")
		if len(valores) != 7 {
			return fmt.Errorf("Formato incorrecto en línea: %s", linea)
		}

		// Eliminar espacios en blanco al inicio y al final de cada valor
		for i := range valores {
			valores[i] = strings.TrimSpace(valores[i])
		}

		// Construir la consulta SQL
		query := fmt.Sprintf("INSERT INTO precios (id, created_at, updated_at, deleted_at, sala_id, precio, moneda) VALUES (%s, %s, %s, %s, %s, %s,'%s')",
			valores[0], valores[1], valores[2], valores[3], valores[4], valores[5], valores[6])

		// Ejecutar la consulta SQL
		DB.DB.Exec(query)

	}

	fmt.Println("Inserciones completadas desde el archivo:", nombreArchivo)
	return nil
}

/*Función que lee los datos desde un archivo txt*/
func InsertarPromociones(nombreArchivo string) error {

	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		return err
	}
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		valores := strings.Split(linea, ",")
		if len(valores) != 7 {
			return fmt.Errorf("Formato incorrecto en línea: %s", linea)
		}

		// Eliminar espacios en blanco al inicio y al final de cada valor
		for i := range valores {
			valores[i] = strings.TrimSpace(valores[i])
		}

		// Construir la consulta SQL
		query := fmt.Sprintf("INSERT INTO promociones (id, created_at, updated_at, deleted_at, titulo, descripcion, url) VALUES (%s, %s, %s, %s, '%s', '%s','%s')",
			valores[0], valores[1], valores[2], valores[3], valores[4], valores[5], valores[6])

		// Ejecutar la consulta SQL
		DB.DB.Exec(query)

	}

	fmt.Println("Inserciones completadas desde el archivo:", nombreArchivo)
	return nil
}
func InsertarDatos() {
	if err := InsertarDatosDesdeArchivo("script/salas.txt"); err != nil {
		fmt.Println("Error al insertar datos desde el archivo:", err)
		return
	}

	if err := InsertarPeliculas("script/peliculas.txt"); err != nil {
		fmt.Println("Error al insertar datos desde el archivo:", err)
		return
	}

	if err := InsertarUsuarios("script/usuario.txt"); err != nil {
		fmt.Println("Error al insertar datos desde el archivo:", err)
		return
	}

	if err := InsertarTarjeta("script/tarjeta.txt"); err != nil {
		fmt.Println("Error al insertar datos desde el archivo:", err)
		return
	}

	InsertarButacas()

	if err := InsertarHorarios("script/horarios.txt"); err != nil {
		fmt.Println("Error al insertar datos desde el archivo:", err)
		return
	}

	if err := InsertarPrecios("script/precios.txt"); err != nil {
		fmt.Println("Error al insertar datos desde el archivo:", err)
		return
	}

	if err := InsertarPromociones("script/promociones.txt"); err != nil {
		fmt.Println("Error al insertar datos desde el archivo:", err)
		return
	}
}
