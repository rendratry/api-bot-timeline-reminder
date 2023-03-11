package service

import (
	"api-bot-timeline-reminder/helper"
	"api-bot-timeline-reminder/model/domain"
	"api-bot-timeline-reminder/model/web"
	"github.com/go-playground/validator/v10"
	"time"
)

type PublishRabbitMQServiceImpl struct {
	Validate *validator.Validate
}

func NewPublishRabbitMQServiceImpl(validate *validator.Validate) *PublishRabbitMQServiceImpl {
	return &PublishRabbitMQServiceImpl{
		Validate: validate,
	}
}

func (service *PublishRabbitMQServiceImpl) PublishDelay(request web.PublishDelayRequest) web.PublishDelayResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	scheduledTime := time.Date(request.Time.Year, time.Month(request.Time.Month), request.Time.Day, request.Time.Hour, request.Time.Minute, 0, 0, time.Local)
	scheduledTimeEpoch := scheduledTime.UnixNano() / int64(time.Millisecond)
	idLog := helper.LogPublishDelay(request.Receiver, scheduledTimeEpoch, "-", request.Platform.Whatsapp, request.Platform.Email, request.Platform.Telegram, false)

	structPublish := web.PublishDelayRequest{IdLog: idLog}

	go helper.PublishRabbitmqDelay(structPublish)
	status := domain.PublishRabbitMQ{
		TimeStatus: string(request.Time.Minute),
	}

	return helper.ToPublishDelayResponse(status)
}

func (service *PublishRabbitMQServiceImpl) SendEmail(request web.SendEmailRequest) web.SendEmailResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	helper.PublishRabbitMQSendEmail(request)

	emailSend := domain.SendEmail{
		Subject: request.Subject,
		Email:   request.Email,
	}

	return helper.ToPublishSendEmailResponse(emailSend)
}
