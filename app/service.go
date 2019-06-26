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
func (s *Service) SendMessage(mailRequest mailprovider.MailRequest) (string, error) {
	var lastError error

	for _, p := range s.providers {
		msg, err := p.SendMail(mailRequest)
		if err == nil {
			log.Printf("Provider %v, wszystko działa \n", p)
			return msg, nil
		}

		lastError = err
		log.Println("Trochę chujnia, bo nie wyszło, provider: %T, error: %s", p, err.Error())
	}

	return "", lastError
}