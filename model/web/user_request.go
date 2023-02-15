package web

type InsertDataMahasiswaRequest struct {
	Nama          string `validate:"required" json:"nama"`
	Email         string `validate:"required,email" json:"email"`
	TahunAngkatan string `validate:"required" json:"tahun_angkatan"`
	Nim           string `validate:"required" json:"nim"`
	Prodi         string `validate:"required" json:"prodi"`
}

type InsertDataStaffRequest struct {
	Nama   string `validate:"required" json:"nama"`
	Nip    string `validate:"required" json:"nip"`
	Alamat string `validate:"required" json:"alamat"`
	Status string `validate:"required" json:"status"`
	Email  string `validate:"required,email" json:"email"`
}
