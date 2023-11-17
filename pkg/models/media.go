package models

type MediaType string

const (
	MediaTypePhoto MediaType = "photo"
	MediaTypeAudio MediaType = "audio"
	MediaTypeGif   MediaType = "gif"
	MediaTypeVideo MediaType = "video"
)

type Media struct {
	ID     string
	Title  string
	Type   MediaType
	Tags   []string
	UserID string
}
