package transport

import (
	"encoding/json"
	"github.com/Nowak90210/hypatos_mail/mail/request"
	"io/ioutil"
	"net/http"

	"github.com/Nowak90210/hypatos_mail/app"
)

type Handler struct {
	service *app.Service
}

func newHandler(s *app.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) sendMail(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var mr request.MailRequest
	if err := json.Unmarshal(body, &mr); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := mr.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.SendMessage(mr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
