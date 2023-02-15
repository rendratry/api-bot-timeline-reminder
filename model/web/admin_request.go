package web

type CreateAdminRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

type LoginAdminRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

type RegisterBotRequest struct {
	Email       string `validate:"required" json:"email"`
	NamaLengkap string `validate:"required" json:"nama_lengkap"`
	NoHp        string `validate:"required" json:"no_hp"`
}

type LoginUserRequest struct {
	Email    string `validate:"required" json:"email"`
	Password string `validate:"required" json:"password"`
}
