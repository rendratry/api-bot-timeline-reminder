package service

import (
	"api-bot-timeline-reminder/model/web"
	"context"
)

type ChatbotService interface {
	GetAllMessages(ctx context.Context, request web.GetChatbotMessagesRequest) []web.GetChatbotMessagesResponse
}
