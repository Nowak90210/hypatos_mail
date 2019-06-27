package mailprovider

import (
	"github.com/pkg/errors"
	"regexp"
)

var (
	ErrInvalidField = errors.New("Invalid field")
	ErrEmptyField   = errors.New("Empty field")
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
		return errors.Wrap(ErrEmptyField, "Field from{email} cannot be empty")
	}

	if isEmailValid(m.From.Email) {
		return nil
	}

	return errors.Wrapf(ErrInvalidField, "Field from{email} must be a mail address, '%s' given", m.From.Email)
}

func (m *MailRequest) validateTo() error {
	if m.To.Email == "" {
		return errors.Wrap(ErrEmptyField, "Field to{email} cannot be empty")
	}

	if isEmailValid(m.To.Email) {
		return nil
	}

	return errors.Wrapf(ErrInvalidField, "Field to{email} must be a mail address, '%s' given", m.To.Email)
}

func (m *MailRequest) validateSubject() error {
	if m.Subject == "" {
		return errors.Wrap(ErrEmptyField, "Subject cannot be empty")
	}

	return nil
}

func (m *MailRequest) validateText() error {
	if m.Text == "" {
		return errors.Wrap(ErrEmptyField, "Text cannot be empty")
	}

	return nil
}

func isEmailValid(email string) bool {
	re := regexp.MustCompile(mailRegEx)
	return re.MatchString(email)
}
