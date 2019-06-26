package transport

import (
	"encoding/json"
	"fmt"
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

	var mail app.Mail
	if err := json.Unmarshal(body, &mail); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := mail.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	msg, id, err := h.service.SendMessage(mail)
	if err != nil {
		w.Write([]byte(err.Error()))

	}

	w.Write([]byte("\n"))
	w.Write([]byte(fmt.Sprintf("MSG: %s \n", msg)))
	w.Write([]byte(fmt.Sprintf("ID: %s \n", id)))
}
