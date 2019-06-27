package app

import (
	"errors"
	mailprovider "github.com/Nowak90210/hypatos_mail/mail_provider"
	"testing"
)

var (
	errorMessage    = "Error Message"
	noErrorProvider = mailprovider.ProviderMock{nil}
	errorProvider   = mailprovider.ProviderMock{errors.New(errorMessage)}
)

func TestSendMessage(t *testing.T) {

	testCases := []struct {
		name        string
		mr          mailprovider.MailRequest
		expectedErr error
		providers   []mailprovider.MailProvider
	}{
		{
			"No Error", mailprovider.MailRequest{}, nil, []mailprovider.MailProvider{noErrorProvider},
		},
		{
			"First with Error", mailprovider.MailRequest{}, nil, []mailprovider.MailProvider{errorProvider, noErrorProvider},
		},
		{
			"Both with Error", mailprovider.MailRequest{}, errors.New(errorMessage), []mailprovider.MailProvider{errorProvider, errorProvider},
		},
	}

	for _, tc := range testCases {
		service := NewService(tc.providers)
		err := service.SendMessage(tc.mr)
		if err != nil && err.Error() != tc.expectedErr.Error() {
			t.Fatalf("%s: expected err = '%s', got err = '%s'", tc.name, tc.expectedErr, err)
		}
	}

}
