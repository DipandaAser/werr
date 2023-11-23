package handlers

import (
	"context"
	"github.com/DipandaAser/werr/internal/auth"
	"github.com/DipandaAser/werr/internal/database"
	"github.com/DipandaAser/werr/pkg/models"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

const (
	// AuthHeader is the header key for the auth token
	AuthHeader = "token"
	// UserFromContextKey is the key to get the user from the context
	UserFromContextKey = "werr_connected_user"
)

func AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("token")

	token, err := auth.GetAuthClient().VerifyIDToken(context.TODO(), tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.NewBadRequestResponse("Invalid token"))
		return
	}

	user, err := database.GetDB().GetUserByID(token.UID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.NewBadRequestResponse("User not found"))
		return
	}

	c.Set(UserFromContextKey, user)
}

func GetUserFromContext(c *gin.Context) models.User {
	return c.MustGet(UserFromContextKey).(models.User)
}

type MaxBytesReader struct {
	rdr        io.ReadCloser
	remaining  int64
	wasAborted bool
	sawEOF     bool
}

func (mbr *MaxBytesReader) Read(p []byte) (n int, err error) {
	toRead := mbr.remaining
	if mbr.remaining == 0 {
		if mbr.sawEOF {
			return 0, models.ErrorFileTooLarge
		}
		// The underlying io.Reader may not return (0, io.EOF)
		// at EOF if the requested size is 0, so read 1 byte
		// instead. The io.Reader docs are a bit ambiguous
		// about the return value of Read when 0 bytes are
		// requested, and {bytes,strings}.Reader gets it wrong
		// too (it returns (0, nil) even at EOF).
		toRead = 1
	}
	if int64(len(p)) > toRead {
		p = p[:toRead]
	}
	n, err = mbr.rdr.Read(p)
	if err == io.EOF {
		mbr.sawEOF = true
	}
	if mbr.remaining == 0 {
		// If we had zero bytes to read remaining (but hadn't seen EOF)
		// and we get a byte here, that means we went over our limit.
		if n > 0 {
			return 0, models.ErrorFileTooLarge
		}
		return 0, err
	}
	mbr.remaining -= int64(n)
	if mbr.remaining < 0 {
		mbr.remaining = 0
	}
	return n, err
}

func (mbr *MaxBytesReader) Close() error {
	return mbr.rdr.Close()
}

// LimiterReader returns a custom io.ReadCloser that limits the size the request
func LimiterReader(body io.ReadCloser, limit int64) *MaxBytesReader {
	return &MaxBytesReader{
		rdr:        body,
		remaining:  limit,
		wasAborted: false,
		sawEOF:     false,
	}
}
