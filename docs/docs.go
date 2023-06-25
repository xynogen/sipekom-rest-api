// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/api": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Root"
                ],
                "summary": "server status.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/absen/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get all Absen",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Absen"
                ],
                "summary": "get all Absen.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/absen/create/{location}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get Absen by location.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Absen"
                ],
                "summary": "create Absen.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "location",
                        "name": "location",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/absen/delete/{id_absen}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete Absen by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Absen"
                ],
                "summary": "delete Absen.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Absen ID",
                        "name": "id_absen",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/absen/get/{id_user}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get Absen by ID User.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Absen"
                ],
                "summary": "get Absen.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID User",
                        "name": "id_user",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/absen/update/{id_absen}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update Absen by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Absen"
                ],
                "summary": "update Absen.",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateAbsenRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Absen ID",
                        "name": "id_absen",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/check": {
            "get": {
                "description": "get validation of the token.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Check"
                ],
                "summary": "encpoint to check token validation.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/elogbook/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get all ELogBook",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ELogBook"
                ],
                "summary": "get all ELogBook.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/elogbook/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "create new ELogBook.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ELogBook"
                ],
                "summary": "create ELogBook.",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateELogBookRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/elogbook/delete/{id_elogbook}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delet ELogBook by ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ELogBook"
                ],
                "summary": "delete ELogBook.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ELogBook ID",
                        "name": "id_elogbook",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/elogbook/get/{id_user}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get ELogBook by id user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ELogBook"
                ],
                "summary": "get ELogBook.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID User",
                        "name": "id_user",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/konsulen/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get all absen",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Konsulen"
                ],
                "summary": "get all Konsulen.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/konsulen/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "create new Konsulen.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Konsulen"
                ],
                "summary": "create Konsulen.",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateKonsulenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/konsulen/get/{id_user}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get Konsulen by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Konsulen"
                ],
                "summary": "get Konsulen.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID User",
                        "name": "id_user",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/konsulen/update/{id_user}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update Konsulen by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Konsulen"
                ],
                "summary": "update Konsulen.",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateKonsulenRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id_user",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "Login and Receive JWT Token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorization"
                ],
                "summary": "Authorization.",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/user/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get all User",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get all User.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/user/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "create new User.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "create User.",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/user/delete/{id_user}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "delete User by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "delete User.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id_user",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/user/get/{id_user}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get User by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "get User.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id_user",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/user/update/{id_elogbook}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update ELogBook by ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ELogBook"
                ],
                "summary": "update ELogBook.",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateELogBookRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "ELogBook ID",
                        "name": "id_elogbook",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/user/update/{id_user}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "update User by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "update User.",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateUserRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id_user",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.CreateELogBookRequest": {
            "type": "object",
            "properties": {
                "deskripsi": {
                    "type": "string"
                },
                "end_time": {
                    "type": "integer"
                },
                "jumlah": {
                    "type": "integer"
                },
                "medical_record": {
                    "type": "string"
                },
                "start_time": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "request.CreateKonsulenRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "spesialis": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.CreateUserRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.LoginRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.UpdateAbsenRequest": {
            "type": "object",
            "properties": {
                "absen": {
                    "type": "integer"
                },
                "absen_flag": {
                    "type": "integer"
                },
                "lokasi": {
                    "type": "string"
                }
            }
        },
        "request.UpdateELogBookRequest": {
            "type": "object",
            "properties": {
                "deskripsi": {
                    "type": "string"
                },
                "end_time": {
                    "type": "integer"
                },
                "jumlah": {
                    "type": "integer"
                },
                "medical_record": {
                    "type": "string"
                },
                "start_time": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "request.UpdateKonsulenRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "spesialis": {
                    "type": "string"
                }
            }
        },
        "request.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "is_activated": {
                    "type": "integer"
                },
                "role": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "API Token",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API Sipekom",
	Description:      "API yang digunakan untuk website SIPEKOM",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
