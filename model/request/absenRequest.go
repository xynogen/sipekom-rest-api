package request

type UpdateAbsenRequest struct {
	Absen     int64  `json:"absen"`
	AbsenFlag uint8  `json:"absen_flag"`
	Lokasi    string `json:"lokasi"`
}
