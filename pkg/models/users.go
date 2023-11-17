package models

import "time"

type User struct {
	ID           string
	JoinedAt     time.Time
	UserName     string
	UserImageUrl string
}
