package request

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Username    string `json:"username"`
	Role        uint8  `json:"role"`
	IsActivated uint8  `json:"is_activated"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     uint8  `json:"role"`
}
