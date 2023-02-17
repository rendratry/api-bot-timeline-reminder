package service

import (
	"api-bot-timeline-reminder/helper"
	"api-bot-timeline-reminder/model/domain"
	"api-bot-timeline-reminder/model/web"
	"api-bot-timeline-reminder/repository"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
)

type ChatbotServiceImpl struct {
	ChatbotRepository repository.ChatbotRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewChatbotService(chatbotRepository repository.ChatbotRepository, DB *sql.DB, validate *validator.Validate) *ChatbotServiceImpl {
	return &ChatbotServiceImpl{
		ChatbotRepository: chatbotRepository,
		DB:                DB,
		Validate:          validate}
}

func (service *ChatbotServiceImpl) GetAllMessages(ctx context.Context, request web.GetChatbotMessagesRequest) []web.GetChatbotMessagesResponse {

	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	getMessages := domain.ChatbotMessages{
		Limit:  request.Limit,
		Offset: request.Offset,
	}

	getAllMessages := service.ChatbotRepository.GetAllMessage(ctx, tx, getMessages)
	return helper.ToGetMessagesResponses(getAllMessages)
}
