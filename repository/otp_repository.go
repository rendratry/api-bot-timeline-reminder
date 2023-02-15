package repository

type OtpRepository interface {
	SendOTP()
	VerifOTP()
}
