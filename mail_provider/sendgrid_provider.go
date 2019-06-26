package mailprovider

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"os"
)

func init() {
	if os.Getenv("SENDGRID_API_KEY") == "" {
		log.Fatalln("Environment variable 'SENDGRID_API_KEY' cannot be emtpy!")
	}
}

type SendGridProvider struct {
	apiKey string
}

func NewSendGridProvider() *SendGridProvider {
	return &SendGridProvider{
		apiKey: os.Getenv("SENDGRID_API_KEY"),
	}
}

func (p *SendGridProvider) SendMail(mr MailRequest) (string, error) {
	message := p.generateMessageFromMailRequest(mr)
	client := sendgrid.NewSendClient(p.apiKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Zrób z tego jakiś ładny message
	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)

	msg := string(response.StatusCode) + ", body: " + response.Body
	return msg, nil
}

func (p *SendGridProvider) generateMessageFromMailRequest(mr MailRequest) *mail.SGMailV3 {
	from := mail.NewEmail(mr.From.Name, mr.From.Email)
	to := mail.NewEmail(mr.To.Name, mr.To.Email)
	message := mail.NewSingleEmail(from, mr.Subject, to, mr.Text, " ")

	return message
}
