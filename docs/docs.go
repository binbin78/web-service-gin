// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "email": "soberkoder@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/albums": {
            "get": {
                "description": "Get details of all albums",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "albums"
                ],
                "summary": "Get details of all albums",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new Album with the input paylod",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "albums"
                ],
                "summary": "Create a new Album",
                "parameters": [
                    {
                        "description": "Album details",
                        "name": "album-model",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Album"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/albums/album/": {
            "put": {
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "albums"
                ],
                "summary": "Add a new albume to the store",
                "parameters": [
                    {
                        "description": "Album Data",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Album"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/albums/{id}": {
            "get": {
                "description": "Get details of album corresponding to the input albumId",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "albums"
                ],
                "summary": "Get details for a given albumId",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the album",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete by album ID",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "album"
                ],
                "summary": "Delete an album",
                "parameters": [
                    {
                        "type": "string",
                        "description": "album ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Album": {
            "type": "object",
            "required": [
                "artist",
                "price",
                "title"
            ],
            "properties": {
                "artist": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API",
	Description:      "This is a sample serice for managing albums",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}