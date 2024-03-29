package request

type CreateELogBookRequest struct {
	IDkonsulen     uint   `json:"id_konsulen"`
	Title          string `json:"title"`
	Jumlah         uint   `json:"jumlah"`
	StartTime      int64  `json:"start_time"`
	EndTime        int64  `json:"end_time"`
	Deskripsi      string `json:"deskripsi"`
	Medical_Record string `json:"medical_record"`
}

type UpdateELogBookRequest struct {
	Title          string `json:"title"`
	Jumlah         uint   `json:"jumlah"`
	StartTime      int64  `json:"start_time"`
	EndTime        int64  `json:"end_time"`
	Deskripsi      string `json:"deskripsi"`
	Medical_Record string `json:"medical_record"`
}
