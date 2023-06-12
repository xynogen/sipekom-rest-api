package request

import "time"

type UpdateAbsenRequest struct {
	Absen     time.Time `json:"absen"`
	AbsenFlag uint8     `json:"absen_flag"`
	Lokasi    string    `json:"lokasi"`
	IDUser    uint      `json:"id_user"`
}
