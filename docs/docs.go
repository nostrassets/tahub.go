// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Alby",
            "url": "https://getalby.com",
            "email": "hello@getalby.com"
        },
        "license": {
            "name": "GNU GPLv3",
            "url": "https://www.gnu.org/licenses/gpl-3.0.en.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v2/admin/users": {
            "put": {
                "description": "Update an account with a new a login, password and activation status. Requires Authorization header with admin token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Update an account",
                "parameters": [
                    {
                        "description": "Update User",
                        "name": "account",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.UpdateUserRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.UpdateUserResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/auth": {
            "post": {
                "description": "Authenticate by pubkey",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Authenticate by pubkey",
                "parameters": [
                    {
                        "description": "Pubkey",
                        "name": "pubkey",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Content",
                        "name": "content",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Refresh Token",
                        "name": "refresh_token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.AuthResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/balance/:asset_id": {
            "get": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Current user's balance in satoshi",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Retrieve balance",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.BalanceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/balances/all": {
            "get": {
                "description": "Retrieve all user balances",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Retrieve all balances",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.BalancesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/create-address": {
            "post": {
                "description": "Get or create address for deposit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Address"
                ],
                "summary": "Get or create address",
                "parameters": [
                    {
                        "description": "Asset ID",
                        "name": "asset_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Amount",
                        "name": "amt",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.AddressResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/invoices": {
            "post": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Returns a new bolt11 invoice",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Invoice"
                ],
                "summary": "Generate a new invoice",
                "parameters": [
                    {
                        "description": "Add Invoice",
                        "name": "invoice",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v2controllers.AddInvoiceRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.AddInvoiceResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/invoices/incoming": {
            "get": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Returns a list of incoming invoices for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Invoice"
                ],
                "summary": "Retrieve incoming invoices",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/v2controllers.Invoice"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/invoices/outgoing": {
            "get": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Returns a list of outgoing payments for a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Invoice"
                ],
                "summary": "Retrieve outgoing payments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/v2controllers.Invoice"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/invoices/{payment_hash}": {
            "get": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Retrieve information about a specific invoice by payment hash",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Invoice"
                ],
                "summary": "Get a specific invoice",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Payment hash",
                        "name": "payment_hash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.Invoice"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/payments/bolt11": {
            "post": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Pay a bolt11 invoice",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payment"
                ],
                "summary": "Pay an invoice",
                "parameters": [
                    {
                        "description": "Invoice to pay",
                        "name": "PayInvoiceRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v2controllers.PayInvoiceRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.PayInvoiceResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/payments/keysend": {
            "post": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Pay a node without an invoice using it's public key",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payment"
                ],
                "summary": "Make a keysend payment",
                "parameters": [
                    {
                        "description": "Invoice to pay",
                        "name": "KeySendRequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v2controllers.KeySendRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.KeySendResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/payments/keysend/multi": {
            "post": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Pay multiple nodes without an invoice using their public key",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payment"
                ],
                "summary": "Make multiple keysend payments",
                "parameters": [
                    {
                        "description": "Invoice to pay",
                        "name": "MultiKeySendRequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v2controllers.MultiKeySendRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.MultiKeySendResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/universe-assets": {
            "get": {
                "description": "Retrieve universe assets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Universe"
                ],
                "summary": "Retrieve universe assets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.UniverseAssetsResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/v2/users": {
            "post": {
                "description": "Create a new account with a pubkey and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Create an account",
                "parameters": [
                    {
                        "description": "Create User",
                        "name": "account",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.CreateUserRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v2controllers.CreateUserResponseBody"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "responses.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "v2controllers.AddInvoiceRequestBody": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer",
                    "minimum": 0
                },
                "description": {
                    "type": "string"
                },
                "description_hash": {
                    "type": "string"
                },
                "ta_asset_id": {
                    "type": "string"
                }
            }
        },
        "v2controllers.AddInvoiceResponseBody": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "string"
                },
                "payment_hash": {
                    "type": "string"
                },
                "payment_request": {
                    "type": "string"
                },
                "ta_asset_id": {
                    "type": "string"
                }
            }
        },
        "v2controllers.AddressResponseBody": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                }
            }
        },
        "v2controllers.AuthResponseBody": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "pubkey": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "v2controllers.BalanceResponse": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "currency": {
                    "type": "string"
                },
                "unit": {
                    "type": "string"
                }
            }
        },
        "v2controllers.BalancesResponse": {
            "type": "object",
            "properties": {
                "balances": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                }
            }
        },
        "v2controllers.CreateUserRequestBody": {
            "type": "object",
            "properties": {
                "pubkey": {
                    "type": "string"
                }
            }
        },
        "v2controllers.CreateUserResponseBody": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "pubkey": {
                    "type": "string"
                }
            }
        },
        "v2controllers.Invoice": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "custom_records": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "array",
                        "items": {
                            "type": "integer"
                        }
                    }
                },
                "description": {
                    "type": "string"
                },
                "description_hash": {
                    "type": "string"
                },
                "destination": {
                    "type": "string"
                },
                "error_message": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "string"
                },
                "fee": {
                    "type": "integer"
                },
                "is_paid": {
                    "type": "boolean"
                },
                "keysend": {
                    "type": "boolean"
                },
                "payment_hash": {
                    "type": "string"
                },
                "payment_preimage": {
                    "type": "string"
                },
                "payment_request": {
                    "type": "string"
                },
                "settled_at": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "v2controllers.KeySendRequestBody": {
            "type": "object",
            "required": [
                "amount",
                "destination",
                "ta_asset_id"
            ],
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "customRecords": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "custom_records": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "destination": {
                    "type": "string"
                },
                "memo": {
                    "type": "string"
                },
                "ta_asset_id": {
                    "type": "string"
                }
            }
        },
        "v2controllers.KeySendResponseBody": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "asset_id": {
                    "type": "integer"
                },
                "custom_records": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "description_hash": {
                    "type": "string"
                },
                "destination": {
                    "type": "string"
                },
                "fee": {
                    "type": "integer"
                },
                "payment_hash": {
                    "type": "string"
                },
                "payment_preimage": {
                    "type": "string"
                }
            }
        },
        "v2controllers.KeySendResult": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/responses.ErrorResponse"
                },
                "keysend": {
                    "$ref": "#/definitions/v2controllers.KeySendResponseBody"
                }
            }
        },
        "v2controllers.MultiKeySendRequestBody": {
            "type": "object",
            "properties": {
                "keysends": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v2controllers.KeySendRequestBody"
                    }
                }
            }
        },
        "v2controllers.MultiKeySendResponseBody": {
            "type": "object",
            "properties": {
                "keysends": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/v2controllers.KeySendResult"
                    }
                }
            }
        },
        "v2controllers.PayInvoiceRequestBody": {
            "type": "object",
            "required": [
                "invoice"
            ],
            "properties": {
                "amount": {
                    "type": "integer",
                    "minimum": 0
                },
                "invoice": {
                    "type": "string"
                },
                "ta_asset_id": {
                    "description": "NOTE: this is defaulting to bitcoin balance",
                    "type": "string",
                    "minLength": 1
                }
            }
        },
        "v2controllers.PayInvoiceResponseBody": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "description_hash": {
                    "type": "string"
                },
                "destination": {
                    "type": "string"
                },
                "fee": {
                    "type": "integer"
                },
                "payment_hash": {
                    "type": "string"
                },
                "payment_preimage": {
                    "type": "string"
                },
                "payment_request": {
                    "type": "string"
                },
                "ta_asset_id": {
                    "type": "string"
                }
            }
        },
        "v2controllers.UniverseAssetsResponseBody": {
            "type": "object",
            "properties": {
                "assets": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                }
            }
        },
        "v2controllers.UpdateUserRequestBody": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "deactivated": {
                    "type": "boolean"
                },
                "deleted": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "pubkey": {
                    "type": "string"
                }
            }
        },
        "v2controllers.UpdateUserResponseBody": {
            "type": "object",
            "properties": {
                "deactivated": {
                    "type": "boolean"
                },
                "deleted": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "pubkey": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "OAuth2Password": {
            "type": "oauth2",
            "flow": "password",
            "tokenUrl": "/v2/auth"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.9.0",
	Host:             "",
	BasePath:         "/v2",
	Schemes:          []string{},
	Title:            "LndHub.go",
	Description:      "Accounting wrapper for the Lightning Network providing separate accounts for end-users.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
