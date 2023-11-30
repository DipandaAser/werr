package handlers

import (
	"context"
	"errors"
	"github.com/DipandaAser/werr/internal/config"
	"github.com/DipandaAser/werr/internal/database"
	"github.com/DipandaAser/werr/internal/storage"
	"github.com/DipandaAser/werr/pkg/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"log"
	"net/http"
)

func UploadMedia(c *gin.Context) {
	user := GetUserFromContext(c)

	db := database.GetDB()
	opts := options.Session().SetDefaultReadConcern(readconcern.Majority())
	sess, err := db.DB.Client().StartSession(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.EndSession(db.Context)

	media := &models.Media{
		Title:  c.Query("title"),
		UserID: user.ID,
		Status: models.MediaStatusDraft,
	}

	_, err = sess.WithTransaction(context.TODO(), func(ctx mongo.SessionContext) (interface{}, error) {
		err := db.CreateMedia(ctx, media)
		if err != nil {
			return nil, err
		}

		limitedReader := LimiterReader(c.Request.Body, config.GetConfig().FILESIZELIMIT)

		defer func() {
			err := limitedReader.Close()
			if err != nil {
				c.Header("connection", "close")
				log.Println(err)
			}
		}()

		err = storage.GetStorage().UploadMedia(limitedReader, user.ID, media.InternalID)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})
	if err != nil {
		if errors.As(err, &models.ErrorFileTooLarge) {
			c.Header("connection", "close")
			c.JSON(http.StatusRequestEntityTooLarge, models.NewBadRequestResponse("The provided file is too large"))
			return
		}
		c.JSON(http.StatusInternalServerError, models.NewBadRequestResponse("An error occurred"))
		return
	}

	c.JSON(http.StatusOK, models.PostMediaResponse{Media: *media})
}
