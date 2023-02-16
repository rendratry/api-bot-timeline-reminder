package repository

import (
	"api-bot-timeline-reminder/model/domain"
	"context"
	"database/sql"
)

type OtpRepository interface {
	SendOtp(ctx context.Context, tx *sql.Tx, otp domain.Otp) domain.Otp
	VerifikasiOtp(ctx context.Context, tx *sql.Tx, otp domain.Otp) (domain.Otp, error)
}
