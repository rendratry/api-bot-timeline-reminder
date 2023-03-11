package web

type PublishDelayResponse struct {
	ScheduledAt string `json:"scheduled_at"`
	Status      string `json:"status"`
}

type SendEmailResponse struct {
	Email  string `json:"email"`
	Status string `json:"status"`
}
