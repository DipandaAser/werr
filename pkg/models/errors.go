package models

import "errors"

var ErrorFileTooLarge = errors.New("file is too large")

type BadRequestResponse struct {
	Error string `json:"error"`
}

func NewBadRequestResponse(err string) BadRequestResponse {
	return BadRequestResponse{
		Error: err,
	}
}
