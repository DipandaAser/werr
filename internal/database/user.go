package database

import (
	"github.com/DipandaAser/werr/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	UserCollectionName = "users"
)

func (db *DB) GetUserByID(id string) (models.User, error) {
	var user models.User
	err := db.DB.Collection(UserCollectionName).
		FindOne(db.Context, bson.M{"_id": id}).
		Decode(&user)
	return user, err
}

func (db *DB) CreateUser(user models.User) error {
	_, err := db.DB.Collection(UserCollectionName).
		InsertOne(db.Context, user)
	return err
}
