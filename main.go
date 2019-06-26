package main

import (
	mailprovider "github.com/Nowak90210/hypatos_mail/mail_provider"
	"net/http"

	"github.com/Nowak90210/hypatos_mail/app"

	"github.com/Nowak90210/hypatos_mail/transport"
)

func main() {
	var providers []mailprovider.MailProvider

	mgProvider := mailprovider.NewMailGunProvider()
	sgProvider := mailprovider.NewSendGridProvider()

	providers = append(providers, mgProvider)
	providers = append(providers, sgProvider)

	service := app.NewService(providers)
	router := transport.InitRouter(service)

	http.ListenAndServe(":8001", router)
}
