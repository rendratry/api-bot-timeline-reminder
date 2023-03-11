package domain

type PublishRabbitMQ struct {
	TimeStatus string
}

type Recipient struct {
	Whatsapp         string
	Email            string
	UsernameTelegram string
	IdTelegram       string
}

type Platform struct {
	Email    bool
	Whatsapp bool
	Telegram bool
}

type SendMsgRabbitMQ struct {
	Message   string
	Platform  Platform
	Recipient Recipient
}

type SendEmail struct {
	Subject string
	Email   string
	Message string
}
