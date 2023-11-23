package database

import (
	"context"
	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type DB struct {
	DB              *mongo.Database
	Context         context.Context
	CancelCtx       context.CancelFunc
	sortIDGenerator *shortid.Shortid
}

var defaultDB *DB

func GetDB() *DB {
	return defaultDB
}

// InitDB initializes the database connection. It panics if the connection fails.
func InitDB(dbURI string, databaseName string) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatal(err)
	}

	defaultDB = &DB{
		DB: client.Database(databaseName),
	}
	defaultDB.Context, defaultDB.CancelCtx = context.WithTimeout(context.Background(), 60*time.Second)
	defaultDB.sortIDGenerator = shortid.GetDefault()
}

func (db *DB) GenerateShortID() (string, error) {
	return db.sortIDGenerator.Generate()
}
