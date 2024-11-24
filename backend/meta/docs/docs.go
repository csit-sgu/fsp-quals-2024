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
        "/countries": {
            "get": {
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Location"
                ],
                "summary": "Get available countries",
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
        "/filter": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Filter"
                ],
                "summary": "Filter event data",
                "parameters": [
                    {
                        "description": "Filter data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.FilterRequest"
                        }
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
        "/localities": {
            "get": {
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Location"
                ],
                "summary": "Get available localities",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Country query parameter",
                        "name": "country",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Region query parameter",
                        "name": "region",
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
        "/notify": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    ""
                ],
                "summary": "Notify service with about changes in database",
                "parameters": [
                    {
                        "description": "Updated events codes",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
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
        "/ping": {
            "get": {
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "misc"
                ],
                "summary": "Ping the service",
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
        "/regions": {
            "get": {
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Location"
                ],
                "summary": "Get available regions",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Required country",
                        "name": "country",
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
        "/sports": {
            "get": {
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sport"
                ],
                "summary": "Get available sports",
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
        "/subscription": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscription"
                ],
                "summary": "Leave an email subscription request",
                "parameters": [
                    {
                        "description": "Subscription Info",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Subscription"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/subscription/confirm": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Subscription"
                ],
                "summary": "Confirm an email subscription request",
                "parameters": [
                    {
                        "description": "Subscription Confirmation",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SubscriptionConfirmation"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Confirmed",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.DateRange": {
            "type": "object",
            "properties": {
                "from": {
                    "type": "string"
                },
                "to": {
                    "type": "string"
                }
            }
        },
        "model.FilterCondition": {
            "type": "object",
            "properties": {
                "additional_info": {
                    "type": "string"
                },
                "age": {
                    "type": "integer"
                },
                "code": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "date_range": {
                    "$ref": "#/definitions/model.DateRange"
                },
                "event_scale": {
                    "type": "string"
                },
                "event_type": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "locality": {
                    "type": "string"
                },
                "region": {
                    "type": "string"
                },
                "sport": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.FilterRequest": {
            "type": "object",
            "properties": {
                "condition": {
                    "$ref": "#/definitions/model.FilterCondition"
                },
                "pagination": {
                    "$ref": "#/definitions/model.Pagination"
                },
                "required_fields": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.Pagination": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                }
            }
        },
        "model.Subscription": {
            "type": "object",
            "properties": {
                "additional_info": {
                    "type": "string"
                },
                "age": {
                    "type": "integer"
                },
                "code": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "end_date": {
                    "type": "string"
                },
                "event_scale": {
                    "type": "string"
                },
                "event_type": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "locality": {
                    "type": "string"
                },
                "region": {
                    "type": "string"
                },
                "sport": {
                    "type": "string"
                },
                "stage": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "model.SubscriptionConfirmation": {
            "type": "object",
            "properties": {
                "confirmation": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "api_key",
            "in": "query"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Backend Service",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
