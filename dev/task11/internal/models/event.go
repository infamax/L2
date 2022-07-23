package models

type Event struct {
	UserID       string `json:"user_id"`
	DateCreated  string `json:"date_created"`
	DateFinished string `json:"date_finished"`
	Description  string `json:"description"`
}
