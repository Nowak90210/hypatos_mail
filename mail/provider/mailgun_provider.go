package provider

import (
	"context"
	"github.com/Nowak90210/hypatos_mail/mail/request"
	"github.com/mailgun/mailgun-go/v3"
	"log"
	"os"
	"time"
)

type MailGunProvider struct {
	domain string
	apiKey string
}

func NewMailGunProvider() *MailGunProvider {
	return &MailGunProvider{
		domain: os.Getenv("MAIL_GUN_DOMAIN"),
		apiKey: os.Getenv("MAIL_GUN_API_KEY"),
	}
}

func init() {
	if os.Getenv("MAIL_GUN_DOMAIN") == "" {
		log.Fatalln("Environment variable 'MAIL_GUN_DOMAIN' cannot be emtpy!")
	}
	if os.Getenv("MAIL_GUN_API_KEY") == "" {
		log.Fatalln("Environment variable 'MAIL_GUN_API_KEY' cannot be emtpy!")
	}
}

func (p *MailGunProvider) SendMail(mr request.MailRequest) error {
	mg := mailgun.NewMailgun(p.domain, p.apiKey)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	message := p.generateMessageFromMailRequest(mr, mg)
	msg, id, err := mg.Send(ctx, message)
	if err != nil {
		log.Println("MailGunProvider/SendMail: ", err)
		return err
	}
	log.Printf("MailGunProvider/SendMail: msg: %s, id: %s ", msg, id)

	return nil
}

func (p *MailGunProvider) generateMessageFromMailRequest(mr request.MailRequest, mg *mailgun.MailgunImpl) *mailgun.Message {
	from := p.generateFrom(mr)
	to := p.generateTo(mr)
	message := mg.NewMessage(from, mr.Subject, mr.Text, to)

	return message
}

func (p *MailGunProvider) generateFrom(mr request.MailRequest) string {
	var from string

	if mr.From.Name != "" {
		from = mr.From.Name + " <" + mr.From.Email + ">"
	} else {
		from = mr.From.Email
	}

	return from
}
func (p *MailGunProvider) generateTo(mr request.MailRequest) string {
	var to string

	if mr.To.Name != "" {
		to = mr.To.Name + " <" + mr.To.Email + ">"
	} else {
		to = mr.To.Email
	}

	return to
}
