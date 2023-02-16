package service

import (
	"api-bot-timeline-reminder/exception"
	"api-bot-timeline-reminder/helper"
	"api-bot-timeline-reminder/model/domain"
	"api-bot-timeline-reminder/model/web"
	"api-bot-timeline-reminder/repository"
	"bytes"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"time"
)

type OtpServiceImpl struct {
	OtpRepository repository.OtpRepository
	DB            *sql.DB
	Validate      *validator.Validate
}

func NewOtpServiceImpl(otpRepository repository.OtpRepository, DB *sql.DB, validate *validator.Validate) *OtpServiceImpl {
	return &OtpServiceImpl{
		OtpRepository: otpRepository,
		DB:            DB,
		Validate:      validate}
}

func (service *OtpServiceImpl) SendOtp(ctx context.Context, request web.OtpRequest) web.OtpResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	gen_otp := helper.EncodeToString(4)

	milis := time.Now().UnixNano() / 1000000

	otp := domain.Otp{
		IdUser: request.IdUser,
		Email:  request.Email,
		NoHp:   request.NoHp,
		Otp:    gen_otp,
		Time:   milis,
	}

	templateEmail := "<!DOCTYPE html>\n<html xmlns:v=\"urn:schemas-microsoft-com:vml\" xmlns:o=\"urn:schemas-microsoft-com:office:office\" lang=\"en\">\n\n<head>\n    <title></title>\n    <meta http-equiv=\"Content-Type\" content=\"text/html; charset=utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n    <!--[if mso]><xml><o:OfficeDocumentSettings><o:PixelsPerInch>96</o:PixelsPerInch><o:AllowPNG/></o:OfficeDocumentSettings></xml><![endif]-->\n    <style>\n        * {\n            box-sizing: border-box;\n        }\n\n        body {\n            margin: 0;\n            padding: 0;\n        }\n\n        a[x-apple-data-detectors] {\n            color: inherit !important;\n            text-decoration: inherit !important;\n        }\n\n        #MessageViewBody a {\n            color: inherit;\n            text-decoration: none;\n        }\n\n        p {\n            line-height: inherit\n        }\n\n        .desktop_hide,\n        .desktop_hide table {\n            mso-hide: all;\n            display: none;\n            max-height: 0px;\n            overflow: hidden;\n        }\n\n        @media (max-width:520px) {\n            .row-content {\n                width: 100% !important;\n            }\n\n            .column .border,\n            .mobile_hide {\n                display: none;\n            }\n\n            table {\n                table-layout: fixed !important;\n            }\n\n            .stack .column {\n                width: 100%;\n                display: block;\n            }\n\n            .mobile_hide {\n                min-height: 0;\n                max-height: 0;\n                max-width: 0;\n                overflow: hidden;\n                font-size: 0px;\n            }\n\n            .desktop_hide,\n            .desktop_hide table {\n                display: table !important;\n                max-height: none !important;\n            }\n        }\n    </style>\n</head>\n\n<body style=\"background-color: #FFFFFF; margin: 0; padding: 0; -webkit-text-size-adjust: none; text-size-adjust: none;\">\n<table class=\"nl-container\" width=\"100%\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt; background-color: #FFFFFF;\">\n    <tbody>\n    <tr>\n        <td>\n            <table class=\"row row-1\" align=\"center\" width=\"100%\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt;\">\n                <tbody>\n                <tr>\n                    <td>\n                        <table class=\"row-content stack\" align=\"center\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt; color: #000000; width: 500px;\" width=\"500\">\n                            <tbody>\n                            <tr>\n                                <td class=\"column column-1\" width=\"100%\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; vertical-align: top; padding-top: 5px; padding-bottom: 5px; border-top: 0px; border-right: 0px; border-bottom: 0px; border-left: 0px;\">\n                                    <table class=\"image_block\" width=\"100%\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt;\">\n                                        <tr>\n                                            <td style=\"width:100%;padding-right:0px;padding-left:0px;\">\n                                                <div align=\"center\" style=\"line-height:10px\"><img src=\"https://firebasestorage.googleapis.com/v0/b/myfin-ktp.appspot.com/o/Logo%2Flauncher_icon.png?alt=media&token=aa7f5759-0a56-43dd-bf44-138dcaa9fe70\" style=\"display: block; height: auto; border: 0; width: 125px; max-width: 100%;\" width=\"125\"></div>\n                                            </td>\n                                        </tr>\n                                    </table>\n                                    <table class=\"divider_block\" width=\"100%\" border=\"0\" cellpadding=\"10\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt;\">\n                                        <tr>\n                                            <td>\n                                                <div align=\"center\">\n                                                    <table border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" width=\"100%\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt;\">\n                                                        <tr>\n                                                            <td class=\"divider_inner\" style=\"font-size: 1px; line-height: 1px; border-top: 1px solid #BBBBBB;\"><span>&#8202;</span></td>\n                                                        </tr>\n                                                    </table>\n                                                </div>\n                                            </td>\n                                        </tr>\n                                    </table>\n                                    <table class=\"heading_block\" width=\"100%\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt;\">\n                                        <tr>\n                                            <td style=\"width:100%;text-align:center;\">\n                                                <h1 style=\"margin: 0; color: #393d47; font-size: 23px; font-family: Arial, Helvetica Neue, Helvetica, sans-serif; line-height: 120%; text-align: center; direction: ltr; font-weight: 700; letter-spacing: normal; margin-top: 0; margin-bottom: 0;\"><span class=\"tinyMce-placeholder\">Verifikasi Email</span></h1>\n                                            </td>\n                                        </tr>\n                                    </table>\n                                </td>\n                            </tr>\n                            </tbody>\n                        </table>\n                    </td>\n                </tr>\n                </tbody>\n            </table>\n            <table class=\"row row-2\" align=\"center\" width=\"100%\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt;\">\n                <tbody>\n                <tr>\n                    <td>\n                        <table class=\"row-content stack\" align=\"center\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt; color: #000000; width: 500px;\" width=\"500\">\n                            <tbody>\n                            <tr>\n                                <td class=\"column column-1\" width=\"100%\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt; font-weight: 400; text-align: left; vertical-align: top; padding-top: 5px; padding-bottom: 5px; border-top: 0px; border-right: 0px; border-bottom: 0px; border-left: 0px;\">\n                                    <table class=\"paragraph_block\" width=\"100%\" border=\"0\" cellpadding=\"10\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;\">\n                                        <tr>\n                                            <td>\n                                                <div style=\"color:#393d47;font-size:14px;font-family:Arial, Helvetica Neue, Helvetica, sans-serif;font-weight:400;line-height:120%;text-align:center;direction:ltr;letter-spacing:0px;\">\n                                                    <p style=\"margin: 0;\">Anda mendaftar ke aplikasi Presensi Marstech, masukkan kode verifikasi untuk melanjutkan berikut kode verifikasi anda</p>\n                                                </div>\n                                            </td>\n                                        </tr>\n                                    </table>\n                                    <table class=\"heading_block\" width=\"100%\" border=\"0\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt;\">\n                                        <tr>\n                                            <td style=\"width:100%;text-align:center;\">\n                                                <h1 style=\"margin: 0; color: #2b66ff; font-size: 29px; font-family: Arial, Helvetica Neue, Helvetica, sans-serif; line-height: 120%; text-align: center; direction: ltr; font-weight: 700; letter-spacing: normal; margin-top: 0; margin-bottom: 0;\"><span class=\"tinyMce-placeholder\">{{.Otp}}</span></h1>\n                                            </td>\n                                        </tr>\n                                    </table>\n                                    <table class=\"paragraph_block\" width=\"100%\" border=\"0\" cellpadding=\"10\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;\">\n                                        <tr>\n                                            <td>\n                                                <div style=\"color:#393d47;font-size:14px;font-family:Arial, Helvetica Neue, Helvetica, sans-serif;font-weight:400;line-height:120%;text-align:center;direction:ltr;letter-spacing:0px;\"></div>\n                                            </td>\n                                        </tr>\n                                    </table>\n                                    <table class=\"paragraph_block\" width=\"100%\" border=\"0\" cellpadding=\"10\" cellspacing=\"0\" role=\"presentation\" style=\"mso-table-lspace: 0pt; mso-table-rspace: 0pt; word-break: break-word;\">\n                                        <tr>\n                                            <td>\n                                                <div style=\"color:#393d47;font-size:14px;font-family:Arial, Helvetica Neue, Helvetica, sans-serif;font-weight:400;line-height:120%;text-align:center;direction:ltr;letter-spacing:0px;\">\n                                                    <p style=\"margin: 0; margin-bottom: 16px;\">Jika Butuh Bantuan silahkan hubungi</p>\n                                                    <p style=\"margin: 0;\"><strong>PT Marstech Global</strong></p>\n                                                </div>\n                                            </td>\n                                        </tr>\n                                    </table>\n                                </td>\n                            </tr>\n                            </tbody>\n                        </table>\n                    </td>\n                </tr>\n                </tbody>\n            </table>\n        </td>\n    </tr>\n    </tbody>\n</table><!-- End -->\n</body>\n\n</html>"

	EmailOTP := helper.EmailTemplate(bytes.Buffer{}, templateEmail, gen_otp)

	helper.SendEmail("Verifikasi Email", request.Email, EmailOTP)

	otp = service.OtpRepository.SendOtp(ctx, tx, otp)

	return helper.ToOtpResponse(otp)
}

func (service *OtpServiceImpl) OtpValidation(ctx context.Context, request web.OtpValidationRequest) web.OtpValidationResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	email := domain.Otp{Email: request.Email, Otp: request.Otp}

	emailvalidation, err := service.OtpRepository.VerifikasiOtp(ctx, tx, email)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	TimeNow := time.Now().UnixNano() / int64(time.Millisecond)

	// Menentukan waktu kadaluarsa sebagai 3 menit dari sekarang
	expire := emailvalidation.Time + int64(3*time.Minute/time.Millisecond)

	var isExpired string
	if TimeNow > expire {
		panic(exception.NewNotFoundError(err.Error()))
	} else {
		isExpired = "success"
	}

	emailvalidation.Email = request.Email
	emailvalidation.Otp = request.Otp
	emailvalidation.Status = isExpired

	return helper.ToOtpValidationResponse(emailvalidation)
}
