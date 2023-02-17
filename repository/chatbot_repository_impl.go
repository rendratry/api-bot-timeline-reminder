package repository

import (
	"api-bot-timeline-reminder/helper"
	"api-bot-timeline-reminder/model/domain"
	"context"
	"database/sql"
)

type ChatbotRepositoryImpl struct {
}

func NewChatbotRepositoryImpl() *ChatbotRepositoryImpl {
	return &ChatbotRepositoryImpl{}
}

func (repository *ChatbotRepositoryImpl) GetAllMessage(ctx context.Context, tx *sql.Tx, messages domain.ChatbotMessages) []domain.ChatbotMessages {
	script := "select id, tag1, tag2, tag3, tag4, tag5, messages from bot_telegram_messages order by id desc limit ? offset ?"
	rows, err := tx.QueryContext(ctx, script, messages.Limit, messages.Offset)
	helper.PanicIfError(err)
	defer rows.Close()

	var newarraymessages []domain.ChatbotMessages
	for rows.Next() {
		newmessages := domain.ChatbotMessages{}
		err := rows.Scan(&newmessages.Id, &newmessages.Tag1, &newmessages.Tag2, &newmessages.Tag3, &newmessages.Tag4, &newmessages.Tag5, &newmessages.Messages)
		helper.PanicIfError(err)
		newarraymessages = append(newarraymessages, newmessages)
	}
	return newarraymessages
}
