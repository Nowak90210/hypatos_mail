package main

import (
	"net/http"

	"github.com/Nowak90210/hypatos_mail/app"

	"github.com/Nowak90210/hypatos_mail/transport"
)

func main() {
	service := app.Service{}
	router := transport.InitRouter(&service)

	http.ListenAndServe(":8001", router)
}
