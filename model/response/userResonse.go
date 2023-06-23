package response

type LoginResponseData struct {
	Username string `json:"username"`
	Role     uint8  `json:"role"`
	ExpireAt int64  `json:"expireAt"`
	IDUser   uint   `json:"id_user"`
}
