package repository

import (
	"api-bot-timeline-reminder/helper"
	"api-bot-timeline-reminder/model/domain"
	"context"
	"database/sql"
	"errors"
)

type OtpRepositoryImpl struct {
}

func NewOtpRepositoryImpl() *OtpRepositoryImpl {
	return &OtpRepositoryImpl{}
}

func (repository *OtpRepositoryImpl) SendOtp(ctx context.Context, tx *sql.Tx, otp domain.Otp) domain.Otp {
	script := "insert into otp(email, no_hp, otp, time_otp) values (?,?,?,?)"
	result, err := tx.ExecContext(ctx, script, otp.Email, otp.NoHp, otp.Otp, otp.Time)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	otp.IdOtp = int(id)

	return otp
}

func (repository *OtpRepositoryImpl) VerifikasiOtp(ctx context.Context, tx *sql.Tx, otp domain.Otp) (domain.Otp, error) {
	script := "select email, time_otp from otp where email = ? and otp = ? limit 1"
	rows, err := tx.QueryContext(ctx, script, otp.Email, otp.Otp)
	otpValidation := domain.Otp{}
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&otpValidation.Email, &otpValidation.Time)
		helper.PanicIfError(err)
		return otpValidation, nil
	} else {
		return otpValidation, errors.New("Kode OTP Salah")
	}
}
