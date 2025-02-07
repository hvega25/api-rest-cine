// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "Prueba rápida para saber si el servidor está en ejecución",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HealthCheck"
                ],
                "summary": "Saber si el servidor esta en ejecución",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/TarjetasDebito": {
            "get": {
                "description": "Obtener tarjetas de débito o crédito para simular una compra",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tarjetas"
                ],
                "summary": "Obtener las tarjetas de débito",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/butaca": {
            "get": {
                "description": "Hace la petición para obtener las butacas mediante la salaID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Butacas"
                ],
                "summary": "Obtiene un item por salaID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "clave foranea a buscar",
                        "name": "sala_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/butacas": {
            "get": {
                "description": "Obtención de todas las butacas en la base de datos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Butacas"
                ],
                "summary": "Obtiene todas las",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/horario": {
            "get": {
                "description": "Hace la petición para obtener un horario mediante la salaID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Horarios"
                ],
                "summary": "Obtiene un item por PeliculaID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "sala a buscar",
                        "name": "pelicula_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/horario/{id}": {
            "get": {
                "description": "Hace la petición para obtener un horario por su id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Horarios"
                ],
                "summary": "Obtiene un horario en especifico",
                "operationId": "Buscar un horario por su id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id del horario a buscar",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/horarios": {
            "get": {
                "description": "Obtener todos los horarios disponibles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Horarios"
                ],
                "summary": "Obtiene todos los horarios disponibles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/precio": {
            "get": {
                "description": "Hace la petición para obtener los precios mediante la salaID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Precios"
                ],
                "summary": "Obtiene precios con la llave foranea SalaID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "llave forenea de sala en precios",
                        "name": "sala_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/precio/{id}": {
            "get": {
                "description": "Hace la petición para obtener una solo precio por su ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Precios"
                ],
                "summary": "Obtiene un precio en especifico",
                "operationId": "Buscar precio por su ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID del precio a buscar",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/precios": {
            "get": {
                "description": "Obtener todos los precios",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Precios"
                ],
                "summary": "Obtener todos los precios",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/promocion/{id}": {
            "get": {
                "description": "Hace la petición para obtener una promocion por su id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Promociones"
                ],
                "summary": "Obtiene una promocion en especifico",
                "operationId": "Buscar una promocion por su id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID de la promocion a buscar",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/promociones": {
            "get": {
                "description": "Obtener todas las promociones disponibles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Promociones"
                ],
                "summary": "Obtener todos los promociones",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/sala/{ID}": {
            "get": {
                "description": "Hace la petición para obtener una sola por su ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Salas"
                ],
                "summary": "Obtiene una sala en especifico",
                "operationId": "Buscar sala por su ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID de la sala a buscar",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/salas": {
            "get": {
                "description": "Obtencion de todas las salas en la base de datos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Salas"
                ],
                "summary": "Obtiene salas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/tarjeta": {
            "get": {
                "description": "Hace la petición para obtener un sola tarjeta por numero",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tarjetas"
                ],
                "summary": "Obtiene un item por su numero de tarjeta",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Numero de la tarjeta a buscar",
                        "name": "cardNumber",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/ticket": {
            "post": {
                "description": "Insertar un nuevo ticket en la base de datos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tickets"
                ],
                "summary": "Insertar un ticket en la base de datos",
                "responses": {
                    "201": {
                        "description": "Creación exitosa",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error en el retorno",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/tickets": {
            "get": {
                "description": "Obtiene todos los tickets ya sea de un usuario registrado o no",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tickets"
                ],
                "summary": "Obtiene todas las tickets",
                "responses": {
                    "200": {
                        "description": "Exitoso",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error en la obtención de tickets",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/usuario": {
            "get": {
                "description": "Hace la petición para obtener un solo usuario por su username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuarios"
                ],
                "summary": "Obtiene un item por su usuario",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username del usuario a buscar",
                        "name": "user_p",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Insertar un nuevo usuario en la base de datos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuarios"
                ],
                "summary": "Insertar un usuario en la base de datos",
                "responses": {
                    "201": {
                        "description": "Retorno exitoso",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error en la obtencion de los datos",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/usuario/{id}": {
            "get": {
                "description": "Hace la petición para obtener un solo usuario por su ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuarios"
                ],
                "summary": "Obtiene un usuario por su correo",
                "operationId": "Busca un usuario por su ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID de la usuario a buscar",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/usuarios": {
            "get": {
                "description": "Obtención de usuarios que ya están registrados en la base de datos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Usuarios"
                ],
                "summary": "Obtención de usuarios",
                "responses": {
                    "201": {
                        "description": "Retorno exitoso",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error en la obtencion de los datos",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/createpelicula": {
            "post": {
                "description": "Insertar un nueva película en la base de datos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Peliculas"
                ],
                "summary": "Insertar una película en la base de datos",
                "responses": {
                    "201": {
                        "description": "Éxito",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/pelicula": {
            "get": {
                "description": "Hace la petición para obtener la lista de películas que estén en cartelera",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Peliculas"
                ],
                "summary": "Obtiene una lista de películas en cartelera",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "columna a comprobar",
                        "name": "estreno",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/peliculas": {
            "get": {
                "description": "Obtencion de todas las peliculas en la base de datos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Peliculas"
                ],
                "summary": "Obtiene peliculas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/peliculas/{id}": {
            "get": {
                "description": "Hace la petición para obtener una sola pelicula por su ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Peliculas"
                ],
                "summary": "Obtiene peliculas",
                "operationId": "Buscar elemento por su ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID de la película a buscar",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:4444",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Cine",
	Description:      "Api rest en GO de un cine",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
