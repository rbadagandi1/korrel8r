// Copyright: This file is part of korrel8r, released under https://github.com/korrel8r/korrel8r/blob/main/LICENSE

// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "url": "https://github.com/korrel8r/korrel8r"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/domains": {
            "get": {
                "tags": [
                    "configuration"
                ],
                "summary": "List all korrel8r domain names.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/graphs/goals": {
            "post": {
                "tags": [
                    "search"
                ],
                "summary": "Create a correlation graph from start objects to goal queries.",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "include rules in graph edges",
                        "name": "withRules",
                        "in": "query"
                    },
                    {
                        "description": "search from start to goal classes",
                        "name": "start",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.GoalsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Graph"
                        }
                    }
                }
            }
        },
        "/graphs/neighbours": {
            "post": {
                "tags": [
                    "search"
                ],
                "summary": "Create a correlation graph of neighbours of a start object to a given depth.",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "include rules in graph edges",
                        "name": "withRules",
                        "in": "query"
                    },
                    {
                        "description": "search from neighbours",
                        "name": "start",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.NeighboursRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Graph"
                        }
                    }
                }
            }
        },
        "/lists/goals": {
            "post": {
                "tags": [
                    "search"
                ],
                "summary": "Generate a list of goal nodes related to a starting point.",
                "parameters": [
                    {
                        "description": "search from start to goal classes",
                        "name": "start",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.GoalsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.Node"
                            }
                        }
                    }
                }
            }
        },
        "/stores": {
            "get": {
                "tags": [
                    "configuration"
                ],
                "summary": "List of all store configurations objects.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/api.StoreConfig"
                            }
                        }
                    }
                }
            }
        },
        "/stores/{domain}": {
            "get": {
                "tags": [
                    "configuration"
                ],
                "summary": "List of all store configurations objects for a specific domain.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "domain\tname",
                        "name": "domain",
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
                                "$ref": "#/definitions/api.StoreConfig"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Edge": {
            "type": "object",
            "properties": {
                "goal": {
                    "description": "Goal is the class name of the goal node.",
                    "type": "string",
                    "example": "class.domain"
                },
                "rules": {
                    "description": "Rules is the set of rules followed along this edge (optional).",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Rule"
                    }
                },
                "start": {
                    "description": "Start is the class name of the start node.",
                    "type": "string"
                }
            }
        },
        "api.GoalsRequest": {
            "description": "Starting point for a goals search.",
            "type": "object",
            "properties": {
                "goals": {
                    "description": "Goal classes for correlation.",
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "class.domain"
                    ]
                },
                "start": {
                    "description": "Start of correlation search.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/api.Start"
                        }
                    ]
                }
            }
        },
        "api.Graph": {
            "description": "Graph resulting from a correlation search.",
            "type": "object",
            "properties": {
                "edges": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Edge"
                    }
                },
                "nodes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/api.Node"
                    }
                }
            }
        },
        "api.NeighboursRequest": {
            "description": "Starting point for a neighbours search.",
            "type": "object",
            "properties": {
                "depth": {
                    "description": "Max depth of neighbours graph.",
                    "type": "integer"
                },
                "start": {
                    "description": "Start of correlation search.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/api.Start"
                        }
                    ]
                }
            }
        },
        "api.Node": {
            "type": "object",
            "properties": {
                "class": {
                    "description": "Class is the full name of the class in \"CLASS.DOMAIN\" form.",
                    "type": "string",
                    "example": "class.domain"
                },
                "count": {
                    "description": "Count of results found for this class, after de-duplication.",
                    "type": "integer"
                },
                "queries": {
                    "description": "Queries yielding results for this class.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/api.Queries"
                        }
                    ],
                    "example": {
                        "querystring": 10
                    }
                }
            }
        },
        "api.Queries": {
            "description": "A set of query strings with counts of results found by the query. Value of -1 means the query was not run so result count is unknown.",
            "type": "object",
            "additionalProperties": {
                "type": "integer"
            }
        },
        "api.Rule": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "Name is an optional descriptive name.",
                    "type": "string"
                },
                "queries": {
                    "description": "Queries generated while following this rule.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/api.Queries"
                        }
                    ],
                    "example": {
                        "querystring": 10
                    }
                }
            }
        },
        "api.Start": {
            "description": "Starting point for correlation.",
            "type": "object",
            "properties": {
                "class": {
                    "description": "Class of starting objects",
                    "type": "string",
                    "example": "class.domain"
                },
                "objects": {
                    "description": "Objects in JSON form",
                    "type": "object"
                },
                "query": {
                    "description": "Queries for starting objects",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "api.StoreConfig": {
            "type": "object",
            "additionalProperties": {
                "type": "string"
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1alpha1",
	Host:             "",
	BasePath:         "/api/v1alpha1",
	Schemes:          []string{},
	Title:            "REST API for korrel8r",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
