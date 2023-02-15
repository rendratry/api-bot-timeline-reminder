package web

type InsertDataMahasiswaResponse struct {
	Nama   string `json:"nama"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

type InsertDataStaffResponse struct {
	Nama   string `json:"nama"`
	Email  string `json:"email"`
	Status string `json:"status"`
}
