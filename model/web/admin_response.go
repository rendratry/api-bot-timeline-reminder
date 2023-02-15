package web

type CreateAdminResponse struct {
	IdUser   string `json:"id_user"`
	Username string `json:"username"`
	Status   string `json:"status"`
}

type LoginAdminResponse struct {
	IdUser   string `json:"id_user"`
	Username string `json:"username"`
	Status   string `json:"status"`
}

type RegisterBotResponse struct {
	Email         string `json:"email"`
	PesanValidasi string `json:"pesan_validasi"`
}

type LoginUserResponse struct {
	IdUser string `json:"id_user"`
	Email  string `json:"username"`
	Role   string `json:"role"`
	Status string `json:"status"`
}
