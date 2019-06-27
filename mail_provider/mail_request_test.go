package mailprovider

import (
	"testing"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		name         string
		mr           MailRequest
		errorMessage string
	}{
		{
			"All Data",
			MailRequest{
				From:    EmailType{Name: "Tomasz Nowak", Email: "tomasz.nowak@excample.com"},
				To:      EmailType{Name: "Tomasz Nowak", Email: "tomasz.nowak@excample.com"},
				Subject: "Subject Content",
				Text:    "Message Content",
			},
			"",
		},
		{
			"Only Required",
			MailRequest{
				From:    EmailType{Email: "tomasz.nowak@excample.com"},
				To:      EmailType{Email: "tomasz.nowak@e-tgn.com"},
				Subject: "Subject Content",
				Text:    "Message Content",
			},
			"",
		},
		{
			"Mail Not valid",
			MailRequest{
				From:    EmailType{Email: "Not email"},
				To:      EmailType{Email: "tomasz.nowak@other-excample.com"},
				Subject: "Subject Content",
				Text:    "Message Content",
			},
			"Field from{email} must be a mail address, 'Not email' given: Invalid field",
		},
		{
			"Mail Empty",
			MailRequest{
				From:    EmailType{Email: ""},
				To:      EmailType{Email: "tomasz.nowak@other-excample.com"},
				Subject: "Subject Content",
				Text:    "Message Content",
			},
			"Field from{email} cannot be empty: Empty field",
		},
		{
			"Subject Empty",
			MailRequest{
				From:    EmailType{Email: "tomasz.nowak@excample.com"},
				To:      EmailType{Email: "tomasz.nowak@other-excample.com"},
				Subject: "",
				Text:    "Message Content",
			},
			"Subject cannot be empty: Empty field",
		},
		{
			"Subject Not Set",
			MailRequest{
				From: EmailType{Email: "tomasz.nowak@excample.com"},
				To:   EmailType{Email: "tomasz.nowak@other-excample.com"},
				Text: "Message Content",
			},
			"Subject cannot be empty: Empty field",
		},
		{
			"Text Empty",
			MailRequest{
				From:    EmailType{Email: "tomasz.nowak@excample.com"},
				To:      EmailType{Email: "tomasz.nowak@other-excample.com"},
				Subject: "Message Content",
				Text:    "Message Content",
			},
			"Subject cannot be empty: Empty field",
		},
	}

	for _, tc := range testCases {
		err := tc.mr.Validate()
		if err != nil {

			if err.Error() != tc.errorMessage {
				t.Fatalf("%s: excepted '%s', got '%s'", tc.name, tc.errorMessage, err)
			}

		}
	}
}
