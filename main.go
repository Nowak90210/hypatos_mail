package main

import (
	"github.com/Nowak90210/hypatos_mail/mail/provider"
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
	var providers []provider.MailProvider

	//mgProvider := provider.NewMailGunProvider()
	sgProvider := provider.NewSendGridProvider()

	//providers = append(providers, mgProvider)
	providers = append(providers, sgProvider)

	return app.NewService(providers)
}
