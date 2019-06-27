package transport

import (
	"bytes"
	"github.com/Nowak90210/hypatos_mail/app"
	mailprovider "github.com/Nowak90210/hypatos_mail/mail_provider"
	_ "github.com/kylelemons/godebug/diff"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// RequestBodies
var (
	validBody = []byte(`{
		"from": {
			"name": "Thomas Nowak",
			"email": "tomasz.nowak@example.com" 
		},
		"subject": "Some Subject",
		"text": "This is a Message Body",
		"to": {
			"name": "Tomasz Nowak",
			"email": "tomasz.grzegorz.nowak@gmail.com"
		}
	}`)
	incompleteBody = []byte(`{
		"subject": "Some Subject",
		"text": "This is a Message Body",
		"to": {
			"name": "Tomasz Nowak",
			"email": "tomasz.grzegorz.nowak@gmail.com"
		}
	}`)
	wrongDataBody = []byte(`Wrong data`)
)
var emptyProviders = []mailprovider.MailProvider{mailprovider.ProviderMock{nil}}
var errorProviders = []mailprovider.MailProvider{mailprovider.ProviderMock{errors.New("Error Message")}}

func TestSendMail(t *testing.T) {

	testCases := []struct {
		name         string
		expectedCode int
		expectedBody string
		jsonBody     []byte
		providers    []mailprovider.MailProvider
	}{
		{"Happy Path", http.StatusAccepted, "", validBody, emptyProviders},
		{"Bad Unmarshall",
			http.StatusBadRequest,
			"Field from{email} cannot be empty: Empty field\n",
			incompleteBody,
			emptyProviders,
		},
		{"Wrong Data",
			http.StatusBadRequest,
			"invalid character 'W' looking for beginning of value\n",
			wrongDataBody,
			emptyProviders,
		},
		{"Happy Path", http.StatusInternalServerError, "Error Message\n", validBody, errorProviders},
	}

	for _, tc := range testCases {
		req, err := http.NewRequest("POST", "localhost:8080/v1/send_mail", bytes.NewBuffer(tc.jsonBody))
		if err != nil {
			t.Fatalf("Faild to create request %s", err)
		}

		rec := httptest.NewRecorder()
		service := app.NewService(tc.providers)
		handler := newHandler(service)

		handler.sendMail(rec, req)
		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != tc.expectedCode {
			t.Fatalf("%s: Expected status code %d, got %d", tc.name, tc.expectedCode, res.StatusCode)
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read body: %v", err)
		}
		if string(b) != tc.expectedBody {
			t.Fatalf("%s: Expected  '%s', got '%s'", tc.name, tc.expectedBody, b)
		}
	}
}
