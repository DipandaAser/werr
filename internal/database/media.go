package database

import (
	"github.com/DipandaAser/werr/pkg/models"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"
)

const (
	MediaCollectionName = "medias"
)

// CreateMedia creates a new media in the database
// On the media parameter, ID, InternalID and CreatedAt will be overwritten
// if the title is empty, the ID will be used as title
func (db *DB) CreateMedia(context mongo.SessionContext, media *models.Media) error {
	var err error
	media.ID, err = db.GenerateShortID()
	if err != nil {
		return err
	}
	if strings.TrimSpace(media.Title) == "" {
		media.Title = media.ID
	}
	if media.Tags == nil {
		media.Tags = make([]string, 0)
	}
	media.InternalID = uuid.Must(uuid.NewRandom()).String()
	media.CreatedAt = time.Now().UTC()
	_, err = db.DB.Collection(MediaCollectionName).InsertOne(context, media)
	return err
}
