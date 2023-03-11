package service

import "api-bot-timeline-reminder/model/web"

type PublishRabbitMQService interface {
	PublishDelay(request web.PublishDelayRequest) web.PublishDelayResponse
	SendEmail(request web.SendEmailRequest) web.SendEmailResponse
}
