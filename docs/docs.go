// Code generated by swaggo/swag. DO NOT EDIT
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
        "/routes": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pickup-locations"
                ],
                "summary": "List of routes between source and each destination.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.RoutesGetResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httputil.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.RoutesGetResp": {
            "type": "object",
            "properties": {
                "routes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.route"
                    }
                },
                "source": {
                    "type": "string"
                }
            }
        },
        "api.route": {
            "type": "object",
            "properties": {
                "destination": {
                    "type": "string"
                },
                "distance": {
                    "type": "number"
                },
                "duration": {
                    "type": "number"
                }
            }
        },
        "httputil.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "httputil.ErrorResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "errors": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/httputil.Error"
                            }
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Pickup Locations API",
	Description:      "Pickup Locations is a service that takes the source and a list of destinations and returns a list of routes between source and each destination.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
