// Package classification Werr API.
//
// Documentation of Werr API.
//
//	Schemes: http, https
//	BasePath: /
//	Version: 1.0.0
//	Host: localhost:7000
//
//	Consumes:
//	- application/json; charset=utf-8
//
//	Produces:
//	- application/json
//
// swagger:meta
package swagger

import "github.com/DipandaAser/werr/pkg/models"

// swagger:response badRequestResponse
type BadRequestResponseWrapper struct {
	// in:body
	Body models.BadRequestResponse
}
