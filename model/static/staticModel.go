package static

const (
	RoleAdmin     uint8 = 0
	RoleKonsulen  uint8 = 1
	RoleMahasiswa uint8 = 2
)

const (
	AbsenCheckIn  uint8 = 0
	AbsenCheckOut uint8 = 1
)

const (
	StatusSuccess string = "success"
	StatusError   string = "error"
)

const (
	KompetensiSenior uint8 = 0
	KompetensiMadya  uint8 = 1
	KompetensiJunior uint8 = 2
)

const (
	NotActivated uint8 = 0
	Activated    uint8 = 1
)

const (
	AccNotApproved uint8 = 0
	AccApproved    uint8 = 1
	AccOnReview    uint8 = 2
)
