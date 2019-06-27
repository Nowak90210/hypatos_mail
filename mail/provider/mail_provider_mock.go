package provider

import "github.com/Nowak90210/hypatos_mail/mail/request"

type ProviderMock struct {
	Err error
}

func (p ProviderMock) SendMail(mailRequest request.MailRequest) error {
	return p.Err
}
