package app

import (
	"errors"
	"github.com/Nowak90210/hypatos_mail/mail/provider"
	"github.com/Nowak90210/hypatos_mail/mail/request"
	"testing"
)

var (
	errorMessage    = "Error Message"
	noErrorProvider = provider.ProviderMock{nil}
	errorProvider   = provider.ProviderMock{errors.New(errorMessage)}
)

func TestSendMessage(t *testing.T) {

	testCases := []struct {
		name        string
		mr          request.MailRequest
		expectedErr error
		providers   []provider.MailProvider
	}{
		{
			"No Error", request.MailRequest{}, nil, []provider.MailProvider{noErrorProvider},
		},
		{
			"First with Error", request.MailRequest{}, nil, []provider.MailProvider{errorProvider, noErrorProvider},
		},
		{
			"Both with Error", request.MailRequest{}, errors.New(errorMessage), []provider.MailProvider{errorProvider, errorProvider},
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
