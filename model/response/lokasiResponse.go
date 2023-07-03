package response

type GetLokasiResponse struct {
	ID     uint   `json:"id_lokasi"`
	Lokasi string `json:"lokasi"`
	URI    string `json:"uri"`
}
