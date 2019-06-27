package app

import (
	mailprovider "github.com/Nowak90210/hypatos_mail/mail_provider"
	"log"
)

type Service struct {
	providers []mailprovider.MailProvider
}

func NewService(prv []mailprovider.MailProvider) *Service {
	return &Service{
		providers: prv,
	}
}

func (s *Service) SendMessage(mailRequest mailprovider.MailRequest) error {
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
