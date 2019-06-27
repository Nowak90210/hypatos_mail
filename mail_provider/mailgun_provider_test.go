package mailprovider

import "testing"

func TestGenerateFrom(t *testing.T) {
	testCases := []struct {
		name     string
		expected string
		mr       MailRequest
	}{
		{
			"Name And Mail",
			"Tomasz Nowak <tomasz.nowak@cokolwiek.com>",
			MailRequest{From: EmailType{Name: "Tomasz Nowak", Email: "tomasz.nowak@cokolwiek.com"}},
		},
		{
			"Only Mail",
			"tomasz.nowak@cokolwiek.com",
			MailRequest{From: EmailType{Email: "tomasz.nowak@cokolwiek.com"}},
		},
	}

	mg := MailGunProvider{}

	for _, tc := range testCases {
		result := mg.generateFrom(tc.mr)
		if result != tc.expected {
			t.Fatalf("%s: expected '%s', got '%s;", tc.name, tc.expected, result)
		}
	}
}

func TestGenerateTo(t *testing.T) {
	testCases := []struct {
		name     string
		expected string
		mr       MailRequest
	}{
		{
			"Name And Mail",
			"Tomasz Nowak <tomasz.nowak@cokolwiek.com>",
			MailRequest{To: EmailType{Name: "Tomasz Nowak", Email: "tomasz.nowak@cokolwiek.com"}},
		},
		{
			"Only Mail",
			"tomasz.nowak@cokolwiek.com",
			MailRequest{To: EmailType{Email: "tomasz.nowak@cokolwiek.com"}},
		},
	}

	mg := MailGunProvider{}

	for _, tc := range testCases {
		result := mg.generateTo(tc.mr)
		if result != tc.expected {
			t.Fatalf("%s: expected '%s', got '%s;", tc.name, tc.expected, result)
		}
	}
}
