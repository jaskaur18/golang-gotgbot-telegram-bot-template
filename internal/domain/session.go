package domain

type Session struct {
	TelegramID int64  `json:"-"`
	Language   string `json:"language"`
}
