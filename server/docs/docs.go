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
        "/login": {
            "post": {
                "description": "Login with a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "username and password",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.loginUserType"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.loginResponseType"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "description": "Logout",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Logout",
                "responses": {
                    "200": {
                        "description": "Logged out successfully"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Logout failed"
                    }
                }
            }
        },
        "/play": {
            "get": {
                "description": "Get current game",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "play"
                ],
                "summary": "Get current game",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ReturnGame"
                        }
                    },
                    "404": {
                        "description": "Game not found"
                    }
                }
            },
            "delete": {
                "description": "Surrend to current game",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "play"
                ],
                "summary": "Surrend to current game",
                "responses": {
                    "201": {
                        "description": "Surrended"
                    },
                    "404": {
                        "description": "Not in a game"
                    }
                }
            }
        },
        "/play/bot/easy": {
            "get": {
                "description": "Create a game against an easy bot",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "play"
                ],
                "summary": "Create a game against an easy bot",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.NewGame"
                        }
                    },
                    "400": {
                        "description": "Not in a game or double not possible"
                    }
                }
            }
        },
        "/play/bot/hard": {
            "get": {
                "description": "Create a game against an hard bot",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "play"
                ],
                "summary": "Create a game against an hard bot",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.NewGame"
                        }
                    },
                    "400": {
                        "description": "Not in a game or double not possible"
                    }
                }
            }
        },
        "/play/bot/medium": {
            "get": {
                "description": "Create a game against an medium bot",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "play"
                ],
                "summary": "Create a game against an medium bot",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.NewGame"
                        }
                    },
                    "400": {
                        "description": "Not in a game or double not possible"
                    }
                }
            }
        },
        "/play/double": {
            "put": {
                "description": "Accept the double",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "play"
                ],
                "summary": "Accept the double",
                "responses": {
                    "201": {
                        "description": "Double accepted"
                    },
                    "400": {
                        "description": "Not in a game or double not possible"
                    }
                }
            },
            "post": {
                "description": "The player want to double",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "play"
                ],
                "summary": "The player want to double",
                "responses": {
                    "201": {
                        "description": "Value of the red dice after the double",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Not in a game or double not possible"
                    }
                }
            },
            "delete": {
                "description": "Refuse the double",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "play"
                ],
                "summary": "Refuse the double",
                "responses": {
                    "201": {
                        "description": "Double refused"
                    },
                    "400": {
                        "description": "Not in a game or can't refuse double"
                    }
                }
            }
        },
        "/play/local": {
            "get": {
                "description": "Create a local game for playing locally in the same device",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "play"
                ],
                "summary": "Create a local game",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.NewGame"
                        }
                    },
                    "400": {
                        "description": "Already in a game"
                    }
                }
            }
        },
        "/play/moves": {
            "get": {
                "description": "Get possible moves for next turn",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "play"
                ],
                "summary": "Get possible moves for next turn",
                "responses": {
                    "200": {
                        "description": "Dice with all possible moves and the ability to double",
                        "schema": {
                            "$ref": "#/definitions/types.FutureTurn"
                        }
                    },
                    "400": {
                        "description": "Not in a game, not your turn or double requested"
                    }
                }
            },
            "post": {
                "description": "Play all the moves for the current turn",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "play"
                ],
                "summary": "Play all the moves for the current turn",
                "parameters": [
                    {
                        "description": "Moves to play",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.Move"
                            }
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Moves played"
                    },
                    "400": {
                        "description": "Moves not legal, not your turn or not in a game"
                    }
                }
            }
        },
        "/play/search": {
            "get": {
                "description": "Start a matchmaking search for a new game",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "play"
                ],
                "summary": "Start a matchmaking search for a new game",
                "responses": {
                    "201": {
                        "description": "Search started"
                    },
                    "400": {
                        "description": "Already searching or in a game"
                    }
                }
            },
            "delete": {
                "description": "Stop a running matchmaking search",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "play"
                ],
                "summary": "Stop a running matchmaking search",
                "responses": {
                    "204": {
                        "description": "Search stopped"
                    },
                    "400": {
                        "description": "Not searching"
                    }
                }
            }
        },
        "/play/tournament": {
            "post": {
                "description": "Create a new tournament",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tournament"
                ],
                "summary": "Create a new tournament",
                "parameters": [
                    {
                        "description": "Tournament object",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.ReturnTournament"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/types.ReturnTournament"
                        }
                    },
                    "400": {
                        "description": "bad data, tournament alredy open"
                    },
                    "500": {
                        "description": "internal server error"
                    }
                }
            }
        },
        "/play/tournament/list": {
            "get": {
                "description": "List all tournaments you can access",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tournament"
                ],
                "summary": "List all tournaments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.TournamentInfo"
                            }
                        }
                    },
                    "500": {
                        "description": "internal server error"
                    }
                }
            }
        },
        "/play/tournament/{tournament_id}": {
            "get": {
                "description": "Get a tournament",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tournament"
                ],
                "summary": "Get a tournament",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tournament ID",
                        "name": "tournament_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ReturnTournament"
                        }
                    },
                    "404": {
                        "description": "tournament not found"
                    }
                }
            },
            "post": {
                "description": "Join a tournament",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tournament"
                ],
                "summary": "Join a tournament",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tournament ID",
                        "name": "tournament_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.ReturnTournament"
                        }
                    },
                    "400": {
                        "description": "alredy in a tournament"
                    },
                    "404": {
                        "description": "tournament not found"
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "user with password",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.customUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "user created",
                        "schema": {
                            "$ref": "#/definitions/types.User"
                        }
                    }
                }
            }
        },
        "/session": {
            "get": {
                "description": "Get auth session",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authentication"
                ],
                "summary": "Get current auth session",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.User"
                        }
                    },
                    "500": {
                        "description": "error"
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.customUser": {
            "type": "object",
            "properties": {
                "firstname": {
                    "type": "string",
                    "example": "giorgio"
                },
                "lastname": {
                    "type": "string",
                    "example": "rossi"
                },
                "mail": {
                    "type": "string",
                    "example": "giorossi@mail.it"
                },
                "password": {
                    "type": "string",
                    "example": "1234"
                },
                "username": {
                    "type": "string",
                    "example": "gio"
                }
            }
        },
        "handler.loginResponseType": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Login successful"
                },
                "user": {
                    "$ref": "#/definitions/handler.loginResponseUser"
                }
            }
        },
        "handler.loginResponseUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "giorossi@mail.it"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "username": {
                    "type": "string",
                    "example": "gio"
                }
            }
        },
        "handler.loginUserType": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "1234"
                },
                "username": {
                    "type": "string",
                    "example": "gio"
                }
            }
        },
        "types.FutureTurn": {
            "type": "object",
            "properties": {
                "can_double": {
                    "description": "True if the player can double the red dice",
                    "type": "boolean"
                },
                "dices": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "possible_moves": {
                    "type": "array",
                    "items": {
                        "type": "array",
                        "items": {
                            "$ref": "#/definitions/types.Move"
                        }
                    }
                }
            }
        },
        "types.LeaderBoardEntry": {
            "type": "object",
            "properties": {
                "lose": {
                    "type": "integer",
                    "example": 1
                },
                "user": {
                    "type": "string",
                    "example": "Giorgio"
                },
                "win": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "types.Move": {
            "type": "object",
            "properties": {
                "from": {
                    "type": "integer",
                    "example": 1
                },
                "to": {
                    "type": "integer",
                    "example": 2
                }
            }
        },
        "types.NewGame": {
            "type": "object",
            "properties": {
                "dices_p1": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "dices_p2": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "game": {
                    "$ref": "#/definitions/types.ReturnGame"
                }
            }
        },
        "types.ReturnGame": {
            "type": "object",
            "properties": {
                "current_player": {
                    "type": "string",
                    "example": "p1"
                },
                "double_owner": {
                    "type": "string",
                    "example": "all"
                },
                "double_value": {
                    "type": "integer",
                    "example": 1
                },
                "elo1": {
                    "type": "integer",
                    "example": 1000
                },
                "elo2": {
                    "type": "integer",
                    "example": 1000
                },
                "end": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                },
                "game_type": {
                    "type": "string",
                    "example": "online"
                },
                "id": {
                    "type": "integer"
                },
                "p1checkers": {
                    "description": "arr[0] is bar",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "p2checkers": {
                    "description": "arr[0] is bar",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "player1": {
                    "description": "Username of the player",
                    "type": "string",
                    "example": "Giorgio"
                },
                "player2": {
                    "type": "string",
                    "example": "Mario"
                },
                "start": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                },
                "status": {
                    "type": "string",
                    "example": "open"
                },
                "tournament": {
                    "type": "integer",
                    "example": 1
                },
                "want_to_double": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "types.ReturnTournament": {
            "type": "object",
            "properties": {
                "allow_users": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "giorgio",
                        "diego",
                        "marco"
                    ]
                },
                "end": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                },
                "games": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.ReturnGame"
                    }
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "leader_board": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.LeaderBoardEntry"
                    }
                },
                "name": {
                    "type": "string",
                    "example": "Tournament name"
                },
                "owner": {
                    "type": "string",
                    "example": "Giorgio"
                },
                "start": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                },
                "status": {
                    "type": "string",
                    "example": "open"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "giorgio",
                        "diego",
                        "marco"
                    ]
                },
                "visibility": {
                    "type": "string",
                    "example": "public"
                }
            }
        },
        "types.TournamentInfo": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "Tournament name"
                },
                "owner": {
                    "type": "string",
                    "example": "Giorgio"
                }
            }
        },
        "types.User": {
            "type": "object",
            "properties": {
                "elo": {
                    "type": "integer",
                    "example": 1000
                },
                "firstname": {
                    "type": "string",
                    "example": "giorgio"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "is_bot": {
                    "type": "boolean",
                    "example": false
                },
                "lastname": {
                    "type": "string",
                    "example": "rossi"
                },
                "mail": {
                    "type": "string",
                    "example": "giorossi@mail.it"
                },
                "username": {
                    "type": "string",
                    "example": "gio"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
