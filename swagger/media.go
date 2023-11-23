package swagger

import (
	"github.com/DipandaAser/werr/pkg/models"
)

// swagger:route POST /media media idOfMediaPost
// Upload a media
// responses:
//	200: postMediaResponse
//	400: badRequestResponse

// swagger:route PUT /media/{id} media idOfMediaUpdate
// Update media information
// responses:
//	200: postMediaResponse
//	400: badRequestResponse

// swagger:route PUT /media/{id}/thumbnail media idOfMediaUpdateThumbnail
// Update media thumbnail, it works for audio and video media only.
// responses:
//	200: postMediaResponse
//	400: badRequestResponse

// swagger:parameters idOfMediaPost
type PostMediaParams struct {
	// desc: A file content to upload
	// in:body
	Body []byte

	// desc: A file name
	// in:query
	Title string `json:"title"`
}

// This media will have the field Status to draft
// swagger:response postMediaResponse
type PostMediaResponseWrapper struct {
	//in:body
	Body models.PostMediaResponse
}
