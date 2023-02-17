package web

type GetChatbotMessagesRequest struct {
	Limit  int    `validate:"required" json:"limit"`
	Offset string `validate:"required" json:"offset"`
}
