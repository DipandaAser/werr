package models

import "time"

type Comment struct {
	ID        string
	CreatedAt time.Time
	Text      string
	UserID    string
	MediaID   string
}

type CommentResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Text      string    `json:"text"`
	UserID    string    `json:"user_id"`
}
