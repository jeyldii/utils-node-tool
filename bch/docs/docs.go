// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-12-21 05:40:55.151174504 +0300 MSK m=+0.024135202

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "This is BUTTON Node API responseModels documentation",
        "title": "Swagger BUTTON Node API",
        "contact": {
            "name": "API Support",
            "email": "nk"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "1.0"
    },
    "host": "localhost:8080",
    "basePath": "/",
    "paths": {
        "/bch/balance/{address}": {
            "get": {
                "description": "return balance of account in BCH for specific node",
                "produces": [
                    "application/json"
                ],
                "summary": "BCH balance of account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "address",
                        "name": "address",
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
                                "$ref": "#/definitions/responses.BalanceResponse"
                            }
                        }
                    }
                }
            }
        },
        "/bch/utxo/{address}": {
            "get": {
                "description": "return UTXO of account",
                "produces": [
                    "application/json"
                ],
                "summary": "BCH UTXO of account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "address",
                        "name": "address",
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
                                "$ref": "#/definitions/responses.UTXOResponse"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "responses.BalanceResponse": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number",
                    "example": 0
                }
            }
        },
        "responses.UTXO": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "amount": {
                    "type": "number"
                },
                "confirmations": {
                    "type": "integer"
                },
                "height": {
                    "type": "integer"
                },
                "satoshis": {
                    "type": "integer"
                },
                "scriptPubKey": {
                    "type": "string"
                },
                "txid": {
                    "type": "string"
                },
                "vout": {
                    "type": "integer"
                }
            }
        },
        "responses.UTXOResponse": {
            "type": "object",
            "properties": {
                "utxo": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/responses.UTXO"
                    }
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}