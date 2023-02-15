package web

type InsertDataMahasiswaRequest struct {
	UUID          string `validate:"required" json:"uuid"`
	Nama          string `validate:"required" json:"nama"`
	Email         string `validate:"required" validate:"email" json:"email"`
	TahunAngkatan string `validate:"required" json:"tahun_angkatan"`
	Nim           string `validate:"required" json:"nim"`
	Prodi         string `validate:"required" json:"prodi"`
}

type InsertDataStaffRequest struct {
	UUID        string `validate:"required" json:"uuid"`
	Nama        string `validate:"required" json:"nama"`
	Nip         string `validate:"required" json:"nip"`
	Alamat      string `validate:"required" json:"alamat"`
	Status      string `validate:"required" json:"status"`
	Email       string `validate:"required" validate:"email" json:"email"`
	NotifStatus string `validate:"required" json:"notif_status"`
}
