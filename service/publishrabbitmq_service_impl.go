package service

import (
	"api-bot-timeline-reminder/helper"
	"api-bot-timeline-reminder/model/domain"
	"api-bot-timeline-reminder/model/web"
	"github.com/go-playground/validator/v10"
	"strconv"
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
	idLog := helper.LogPublishDelay(request.Receiver, request.FromApp, scheduledTimeEpoch, "-", request.Platform.Whatsapp, request.Platform.Email, request.Platform.Telegram, false)

	structPublish := web.PublishDelayRequest{
		IdLog:        idLog,
		FromApp:      request.FromApp,
		Receiver:     request.Receiver,
		Message:      request.Message,
		EmailSubject: request.EmailSubject,
		EmailMessage: request.EmailMessage,
		Platform: web.Platform{
			Whatsapp: request.Platform.Whatsapp,
			Telegram: request.Platform.Telegram,
			Email:    request.Platform.Email,
		},
		Time: web.TimeDelay{
			Year:   request.Time.Year,
			Month:  request.Time.Month,
			Day:    request.Time.Day,
			Hour:   request.Time.Hour,
			Minute: request.Time.Minute,
		},
	}

	go helper.PublishRabbitmqDelay(structPublish)
	status := domain.PublishRabbitMQ{
		TimeStatus: strconv.FormatInt(scheduledTimeEpoch, 10),
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
