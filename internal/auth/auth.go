package auth

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
	"log"
)

var authClient *auth.Client

func InitAuthClient(firebaseCredFile string) {
	opt := option.WithCredentialsFile(firebaseCredFile)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to create Firebase app: %v", err)
	}

	authClient, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Failed to create Firebase auth client: %v", err)
	}
}

func GetAuthClient() *auth.Client {
	return authClient
}
