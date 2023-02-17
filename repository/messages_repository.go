package repository

import (
	"context"
	"github.com/streadway/amqp"
)

type MessagesRepository interface {
	SendMessages(ctx context.Context, connection *amqp.Connection)
}
