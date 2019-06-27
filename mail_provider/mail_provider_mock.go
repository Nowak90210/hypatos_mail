package mailprovider

type ProviderMock struct {
	Err error
}

func (p ProviderMock) SendMail(mailRequest MailRequest) error {
	return p.Err
}
