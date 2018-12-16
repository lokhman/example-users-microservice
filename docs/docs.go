// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-12-15 23:45:37.3898884 +0000 STD m=+0.057482001

package docs

import (
	"bytes"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is the example microservice for manipulating users.",
        "title": "Example Users Microservice",
        "contact": {
            "name": "Alex Lokhman",
            "url": "https://github.com/lokhman",
            "email": "alex.lokhman@gmail.com"
        },
        "license": {},
        "version": "0.1"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/users": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List users",
                "parameters": [
                    {
                        "maxLength": 2,
                        "minLength": 2,
                        "type": "string",
                        "description": "User country",
                        "name": "country",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.User"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create new user",
                "parameters": [
                    {
                        "description": "New user details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.UserInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.HTTPError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.HTTPError"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "View user details",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "User ID",
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
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update user by ID",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "New user details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/api.UserInput"
                        }
                    }
                ],
                "responses": {
                    "204": {},
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.HTTPError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete user by ID",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {},
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/common.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.UserInput": {
            "type": "object",
            "required": [
                "country",
                "email",
                "first_name",
                "last_name",
                "nickname",
                "password"
            ],
            "properties": {
                "country": {
                    "type": "string",
                    "example": "RU"
                },
                "email": {
                    "type": "string",
                    "example": "alex.lokhman@gmail.com"
                },
                "first_name": {
                    "type": "string",
                    "example": "Alex"
                },
                "last_name": {
                    "type": "string",
                    "example": "Lokhman"
                },
                "nickname": {
                    "type": "string",
                    "example": "VisioN"
                },
                "password": {
                    "type": "string",
                    "example": "MyPassword"
                }
            }
        },
        "common.HTTPError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Error message"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "country": {
                    "type": "string",
                    "example": "RU"
                },
                "email": {
                    "type": "string",
                    "example": "alex.lokhman@gmail.com"
                },
                "first_name": {
                    "type": "string",
                    "example": "Alex"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "last_name": {
                    "type": "string",
                    "example": "Lokhman"
                },
                "nickname": {
                    "type": "string",
                    "example": "VisioN"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo swaggerInfo

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}