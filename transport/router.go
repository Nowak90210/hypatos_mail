package transport

import (
	"net/http"

	"github.com/Nowak90210/hypatos_mail/app"
	"github.com/go-chi/chi"
)

func InitRouter(s *app.Service) http.Handler {
	h := newHandler(s)
	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		r.Post("/send_mail", h.sendMail)
	})

	return r
}
