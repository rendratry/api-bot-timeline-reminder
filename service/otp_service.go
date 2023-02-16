package service

import (
	"api-bot-timeline-reminder/model/web"
	"context"
)

type OtpService interface {
	SendOtp(ctx context.Context, request web.OtpRequest) web.OtpResponse
	OtpValidation(ctx context.Context, request web.OtpValidationRequest) web.OtpValidationResponse
}
