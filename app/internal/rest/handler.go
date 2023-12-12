package rest

import "github.com/julienschmidt/httprouter"

const (
	MimeTypeApplicationZip  = "application/zip"
	MimeTypeApplicationJSON = "application/json"
	MeaderContentType       = "Content-Type"
)

type Handler interface {
	Register(router *httprouter.Router)
}
