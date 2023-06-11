package jsonmodel

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginOutput struct {
	Username string `json:"username"`
	Level    uint8  `json:"level"`
	ExpireAt int64  `json:"expireAt"`
}

type UpdateUserInput struct {
	Username string `json:"username"`
	Level    uint8  `json:"level"`
}
