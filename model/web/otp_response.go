package web

type OtpResponse struct {
	Email  string `json:"email"`
	NoHp   string `json:"no_hp"`
	Status string `json:"status"`
}

type OtpValidationResponse struct {
	Email  string `json:"email"`
	Otp    string `json:"otp"`
	Status string `json:"status"`
}
