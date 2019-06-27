package app

import (
	"github.com/Nowak90210/hypatos_mail/mail/provider"
	"github.com/Nowak90210/hypatos_mail/mail/request"
	"log"
)

type Service struct {
	providers []provider.MailProvider
}

func NewService(prv []provider.MailProvider) *Service {
	return &Service{
		providers: prv,
	}
}

func (s *Service) SendMessage(mailRequest request.MailRequest) error {
	var lastError error

	for _, p := range s.providers {
		err := p.SendMail(mailRequest)
		if err == nil {
			return nil
		}

		lastError = err
		log.Printf("app/service/SendMessage, provider: %T, err: %s ", p, err.Error())
	}

	return lastError
}
