// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Roman Calculator allows you to evaluate mathematical expressions with roman numbers (I, IV, V, X, L, C, M...). Expressions are evaluated according to BODMAS:\n - Brackets ()\n - Order ^\n - Multiplication *\n - Addition +\n - Subtraction -",
    "title": "Roman Calculator",
    "version": "1.0.0"
  },
  "paths": {
    "/calc": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "calc"
        ],
        "summary": "Compute the result of an expression",
        "operationId": "calcExpr",
        "parameters": [
          {
            "type": "string",
            "description": "Roman expression to calculate, e.g. \"I*(II+III)\"",
            "name": "expr",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Result of computing the expression",
            "schema": {
              "$ref": "#/definitions/Result"
            }
          },
          "400": {
            "description": "Invalid request"
          }
        }
      }
    }
  },
  "definitions": {
    "Expr": {
      "type": "string",
      "example": "I*(II+III)"
    },
    "Result": {
      "type": "object",
      "properties": {
        "errs": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "example": []
        },
        "expr": {
          "type": "string",
          "example": "I*(II+III)"
        },
        "result": {
          "type": "string",
          "example": "V"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Roman Calculator allows you to evaluate mathematical expressions with roman numbers (I, IV, V, X, L, C, M...). Expressions are evaluated according to BODMAS:\n - Brackets ()\n - Order ^\n - Multiplication *\n - Addition +\n - Subtraction -",
    "title": "Roman Calculator",
    "version": "1.0.0"
  },
  "paths": {
    "/calc": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "calc"
        ],
        "summary": "Compute the result of an expression",
        "operationId": "calcExpr",
        "parameters": [
          {
            "type": "string",
            "description": "Roman expression to calculate, e.g. \"I*(II+III)\"",
            "name": "expr",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Result of computing the expression",
            "schema": {
              "$ref": "#/definitions/Result"
            }
          },
          "400": {
            "description": "Invalid request"
          }
        }
      }
    }
  },
  "definitions": {
    "Expr": {
      "type": "string",
      "example": "I*(II+III)"
    },
    "Result": {
      "type": "object",
      "properties": {
        "errs": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "example": []
        },
        "expr": {
          "type": "string",
          "example": "I*(II+III)"
        },
        "result": {
          "type": "string",
          "example": "V"
        }
      }
    }
  }
}`))
}
