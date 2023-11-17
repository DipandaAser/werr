package handlers

import (
	"github.com/DipandaAser/werr/internal/auth"
	"github.com/gin-gonic/gin"
)

func UploadMedia(c *gin.Context) {
	auth.GetAuthClient()
}
