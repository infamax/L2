package models

type Event struct {
	UserID       int    `json:"user_id"`
	EventID      int    `json:"event_id"`
	Title        string `json:"title"`
	DateCreated  string `json:"date_created"`
	DateFinished string `json:"date_finished"`
	Description  string `json:"description"`
	Done         bool   `json:"done"`
}
