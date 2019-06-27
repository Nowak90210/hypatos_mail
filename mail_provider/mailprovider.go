package mailprovider

type MailProvider interface {
	SendMail(MailRequest) error
}
