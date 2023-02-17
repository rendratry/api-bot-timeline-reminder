package web

type GetChatbotMessagesRequest struct {
	Limit  int `validate:"required" json:"limit"`
	Offset int `validate:"required" json:"offset"`
}
