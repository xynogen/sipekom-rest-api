package request

type CreateKonsulenRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Spesialis string `json:"spesialis"`
}

type UpdateKonsulenRequest struct {
	Name      string `json:"name"`
	Spesialis string `json:"spesialis"`
}
