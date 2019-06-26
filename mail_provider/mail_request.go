package mailprovider

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidField = errors.New("invalid field")
	ErrEmptyField   = errors.New("empty field")
	mailRegEx       = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)

type MailRequest struct {
	From    EmailType
	Subject string `json:"subject"`
	Text    string `json:"text"`
	To      EmailType
}

type EmailType struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (m *MailRequest) Validate() error {
	if err := m.validateFrom(); err != nil {
		return err
	}
	if err := m.validateTo(); err != nil {
		return err
	}
	if err := m.validateSubject(); err != nil {
		return err
	}
	if err := m.validateText(); err != nil {
		return err
	}

	return nil
}

func (m *MailRequest) validateFrom() error {
	if m.From.Email == "" {
		return ErrEmptyField
	}

	if isEmailValid(m.From.Email) {
		return nil
	}

	return ErrInvalidField
}

func (m *MailRequest) validateTo() error {
	if m.To.Email == "" {
		return ErrEmptyField
	}

	if isEmailValid(m.To.Email) {
		return nil
	}

	return ErrInvalidField
}

func (m *MailRequest) validateSubject() error {
	if m.Subject == "" {
		return ErrEmptyField
	}

	return nil
}

func (m *MailRequest) validateText() error {
	if m.Text == "" {
		return ErrEmptyField
	}

	return nil
}

func isEmailValid(email string) bool {
	re := regexp.MustCompile(mailRegEx)
	return re.MatchString(email)
}
