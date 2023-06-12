package response

type LoginResponseData struct {
	Username string `json:"username"`
	Level    uint8  `json:"level"`
	ExpireAt int64  `json:"expireAt"`
}
