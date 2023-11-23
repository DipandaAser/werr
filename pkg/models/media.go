package models

import "time"

type MediaType string
type MediaStatus string

const (
	MediaTypePhoto MediaType = "photo"
	MediaTypeAudio MediaType = "audio"
	MediaTypeGif   MediaType = "gif"
	MediaTypeVideo MediaType = "video"

	MediaStatusDraft      = "draft"
	MediaStatusProcessing = "processing"
	MediaStatusCompleted  = "completed"
)

// Media is the representation of a media in the database
type Media struct {
	ID         string      `json:"id" bson:"_id"`
	InternalID string      `json:"-" bson:"internal_id"`
	CreatedAt  time.Time   `json:"created_at" bson:"created_at"`
	Title      string      `json:"title" bson:"title"`
	Status     MediaStatus `json:"status" bson:"status"`
	Type       MediaType   `json:"type" bson:"type"`
	Tags       []string    `json:"tags" bson:"tags"`
	UserID     string      `json:"user_id" bson:"user_id"`
}

// PostMediaResponse is the response shown when a media is created or updated
type PostMediaResponse struct {
	Media
}
