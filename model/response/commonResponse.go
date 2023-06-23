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
	IDUser   uint   `json:"id_user"`
	Username string `json:"username"`
	Role     uint8  `json:"role"`
	Exp      int64  `json:"exp"`
}
