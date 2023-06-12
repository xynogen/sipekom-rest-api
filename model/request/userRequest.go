package request

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Username    string `json:"username"`
	Level       uint8  `json:"level"`
	IsActivated uint8  `json:"is_activated"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Level    uint8  `json:"level"`
}
