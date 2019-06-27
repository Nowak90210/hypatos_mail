package main

import (
	mailprovider "github.com/Nowak90210/hypatos_mail/mail_provider"
	"log"
	"net/http"

	"github.com/Nowak90210/hypatos_mail/app"

	"github.com/Nowak90210/hypatos_mail/transport"
)

func main() {
	service := initService()
	router := transport.InitRouter(service)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Server stopped %s", err)
	}
}

func initService() *app.Service {
	var providers []mailprovider.MailProvider

	mgProvider := mailprovider.NewMailGunProvider()
	sgProvider := mailprovider.NewSendGridProvider()

	providers = append(providers, mgProvider)
	providers = append(providers, sgProvider)

	return app.NewService(providers)
}
