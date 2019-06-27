package provider

import "github.com/Nowak90210/hypatos_mail/mail/request"

type MailProvider interface {
	SendMail(request.MailRequest) error
}
