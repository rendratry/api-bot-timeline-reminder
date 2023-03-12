package web

type TimeDelay struct {
	Year   int `validate:"required" json:"year"`
	Month  int `validate:"required" json:"month"`
	Day    int `validate:"required" json:"day"`
	Hour   int `validate:"required" json:"hour"`
	Minute int `validate:"required" json:"minute"`
}

type Platform struct {
	Whatsapp bool `json:"whatsapp"`
	Telegram bool `json:"telegram"`
	Email    bool `json:"email"`
}

type PublishDelayRequest struct {
	IdLog        int
	FromApp      string    `validate:"required"`
	Receiver     string    `validate:"required" json:"receiver"`
	Message      string    `validate:"required" json:"message"`
	EmailSubject string    `validate:"required" json:"email_subject"`
	EmailMessage string    `validate:"required" json:"email_message"`
	Platform     Platform  `validate:"required" json:"platform"`
	Time         TimeDelay `validate:"required" json:"time"`
}

type SendEmailRequest struct {
	FromApp string `validate:"required"`
	Subject string `validate:"required" json:"subject""`
	Email   string `validate:"required" json:"email"`
	Message string `validate:"required" json:"message"`
}
