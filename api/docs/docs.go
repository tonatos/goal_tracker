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
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/goal": {
            "get": {
                "description": "Get goal by :id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Goals"
                ],
                "summary": "Goals Item",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Goal ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/utils.JSONResult"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "data": {
                                                "type": "array",
                                                "items": {
                                                    "$ref": "#/definitions/response.ResponesGoal"
                                                }
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create goal",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Goals"
                ],
                "summary": "Goal Create",
                "parameters": [
                    {
                        "description": "Goal object for create",
                        "name": "goal",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RequestCreateGoal"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.ResponesGoal"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/goal/:goal/contribution": {
            "get": {
                "description": "Get contributions list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contribution"
                ],
                "summary": "Contribution List",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Goal ID for contribution",
                        "name": "goal",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/utils.JSONResult"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "data": {
                                                "type": "array",
                                                "items": {
                                                    "$ref": "#/definitions/response.ResponesContribution"
                                                }
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "description": "Create contribuitions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contribution"
                ],
                "summary": "Contribution Create",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Goal ID for contribution",
                        "name": "goal",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Contribution object for create",
                        "name": "contribuition",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RequestCreateContribution"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.ResponesContribution"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/goal/:goal/contribution/:id": {
            "put": {
                "description": "Update contributions by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contribution"
                ],
                "summary": "Contribution Update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Goal ID for contribution",
                        "name": "goal",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Contribution ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Contribution` + "`" + `s fields for update",
                        "name": "contribution",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RequestUpdateContribution"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.ResponesContribution"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete contribution by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contribution"
                ],
                "summary": "Contribution Delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Goal ID for contribution",
                        "name": "goal",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Contribution ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        },
        "/v1/goal/:id": {
            "get": {
                "description": "Get goals list",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Goals"
                ],
                "summary": "Goal List",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.ResponesGoal"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            },
            "put": {
                "description": "Update goal by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Goals"
                ],
                "summary": "Goal Update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Goal ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Goal` + "`" + `s fields for update",
                        "name": "goal",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RequestUpdateGoal"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.ResponesGoal"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete goal by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Goals"
                ],
                "summary": "Goal Delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Goal ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.RequestCreateContribution": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                }
            }
        },
        "request.RequestCreateGoal": {
            "type": "object",
            "properties": {
                "goal_amount": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "target_date": {
                    "type": "string"
                }
            }
        },
        "request.RequestUpdateContribution": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                }
            }
        },
        "request.RequestUpdateGoal": {
            "type": "object",
            "properties": {
                "goal_amount": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "target_date": {
                    "type": "string"
                }
            }
        },
        "response.ResponesContribution": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "response.ResponesGoal": {
            "type": "object",
            "properties": {
                "ads_by_amount": {
                    "type": "integer"
                },
                "catalog_url": {
                    "type": "string"
                },
                "goal_amount": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "target_date": {
                    "type": "string"
                }
            }
        },
        "utils.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        },
        "utils.JSONResult": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "localhost:3000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "GoalTracker API",
	Description:      "API Service of simple app for tracking your widescale goals",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
