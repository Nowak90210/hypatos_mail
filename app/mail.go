package app

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidField = errors.New("invalid field")
	ErrEmptyField   = errors.New("empty field")
)

type Mail struct {
	From    string `json:"from"`
	Subject string `json:"subject"`
	Text    string `json:"text"`
	To      string `json:"to"`
}

func NewMail(from, subject, text, to string) *Mail {
	return &Mail{
		From:    from,
		Subject: subject,
		Text:    text,
		To:      to,
	}
}

func (m *Mail) Validate() error {
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

func (m *Mail) validateFrom() error {
	if m.From == "" {
		return ErrEmptyField
	}

	if isEmailValid(m.From) {
		return nil
	}

	return ErrInvalidField
}

func (m *Mail) validateTo() error {
	if m.To == "" {
		return ErrEmptyField
	}

	if isEmailValid(m.To) {
		return nil
	}

	return ErrInvalidField
}

func (m *Mail) validateSubject() error {
	if m.Subject == "" {
		return ErrEmptyField
	}

	return nil
}

func (m *Mail) validateText() error {
	if m.Text == "" {
		return ErrEmptyField
	}

	return nil
}

func isEmailValid(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}
