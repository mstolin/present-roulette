package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/mstolin/present-roulette/gmail-adapter/mail"
	"github.com/mstolin/present-roulette/utils/httpErrors"
)

var smtpClientInstance mail.SMTPClient

func NewHandler(smtpClient mail.SMTPClient) http.Handler {
	smtpClientInstance = smtpClient
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.NotFound(notFoundHandler)
	router.Route("/mail", mailHandler)
	return router
}

func notFoundHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(404)
	render.Render(writer, request, &httpErrors.ErrNotFound)
}
