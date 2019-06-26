package app

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v3"
)

var (
	MailGunDomain = os.Getenv("MAIL_GUN_DOMAIN")
	MailGunAPIKey = os.Getenv("MAIL_GUN_API_KEY")
)

//test
func init() {
	fmt.Println("MailGunDomain: ", MailGunDomain)
	fmt.Println("MailGunAPIKey: ", MailGunAPIKey)
}

type Service struct {
	// mail Mail
}

func (s *Service) SendMessage(mail Mail) (string, string, error) {

	mg := mailgun.NewMailgun(MailGunDomain, MailGunAPIKey)

	m := mg.NewMessage(
		"Testowy User <"+mail.From+">",
		mail.Subject,
		mail.Text,
		mail.To,
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	res, err := mg.GetDomain(ctx, MailGunDomain)
	fmt.Println("Res: ", res)
	fmt.Println("Err: ", err)

	return mg.Send(ctx, m)
}
