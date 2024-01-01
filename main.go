package main

import (
	"context"
	"errors"
	"github.com/DipandaAser/werr/internal/auth"
	"github.com/DipandaAser/werr/internal/config"
	"github.com/DipandaAser/werr/internal/database"
	"github.com/DipandaAser/werr/internal/handlers"
	"github.com/DipandaAser/werr/internal/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.LoadConfiguration()
	database.InitDB(config.GetConfig().MONGO_URI, config.GetConfig().DATABASE_NAME)
	storage.InitStorage(storage.StorageConfig{
		AccessKey: config.GetConfig().S3_ACCESS_KEY,
		SecretKey: config.GetConfig().S3_SECRET_KEY,
		Endpoint:  config.GetConfig().S3_ENDPOINT,
		Region:    config.GetConfig().S3_REGION,
		Bucket:    config.GetConfig().S3_BUCKET,
	})
	auth.InitAuthClient(config.GetConfig().FIREBASE_CREDENTIALS_FILE)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Token"},
		ExposeHeaders:    []string{},
		AllowCredentials: true,
		AllowOrigins:     config.GetConfig().ALLOWED_ORIGINS,
	}))

	//Health check endpoint
	router.GET("/health", handlers.Health)

	// Media endpoints
	router.POST("/media", handlers.AuthMiddleware, handlers.UploadMedia)

	srv := &http.Server{
		Addr:    ":7000",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server Panic: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds to finish processing pending requests.
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server started 5sec remaining ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown with error:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}

	log.Println("Server exiting")
}
