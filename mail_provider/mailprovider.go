package mailprovider

type MailProvider interface {
	SendMail(MailRequest) (string, error)
}
