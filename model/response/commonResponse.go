package response

import "github.com/golang-jwt/jwt/v4"

type TokenResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Token   string      `json:"token"`
	Data    interface{} `json:"data"`
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Claims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Level    uint8  `json:"level"`
	Exp      int64  `json:"exp"`
}
