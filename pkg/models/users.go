package models

import "time"

type User struct {
	ID           string    `json:"id" bson:"_id"`
	JoinedAt     time.Time `json:"joined_at" bson:"joined_at"`
	UserName     string    `json:"userName" bson:"userName"`
	UserImageUrl string    `json:"user_image_url" bson:"user_image_url"`
}
