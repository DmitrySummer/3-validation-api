package verifi

import (
	"3-validation-api/config"
	"net/http"
	"net/smtp"
)

type VerifiHandler struct {
	*config.Config
}

func NewAuthHandler(router *http.ServeMux, deps VerifiHandler) {
	handler := &VerifiHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("POST /verify/{hash}", handler.Verify())
}

func (handler *VerifiHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e.Send("smtp.gmail.com:587", smtp.PlainAuth())
	}
}

func (handler *VerifiHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
