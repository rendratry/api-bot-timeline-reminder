package repository

import (
	"api-bot-timeline-reminder/model/domain"
	"context"
	"database/sql"
)

type ChatbotRepository interface {
	GetAllMessage(ctx context.Context, tx *sql.Tx, messages domain.ChatbotMessages) []domain.ChatbotMessages
}
