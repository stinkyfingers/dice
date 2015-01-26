package email

import (
	"bytes"
	// "errors"
	"net/smtp"
	"os"
)

func getAuth() (smtp.Auth, error) {
	if host := os.Getenv("EMAIL_HOST"); host != "" {
		password := os.Getenv("EMAIL_PASSWORD")
		username := os.Getenv("EMAIL_USERNAME")
		return smtp.PlainAuth("", username, password, host), nil

	}
	return smtp.PlainAuth("", "", "", ""), nil
}

func Send(to []string, subject string, doc bytes.Buffer) error {
	addr := "smtp.gmail.com:587"
	from := "john_shenk@hotmail.com"
	if len(to) == 0 {
		to = []string{"john_shenk@hotmail.com"}
	}

	a, err := getAuth()
	if err != nil {
		return err
	}
	err = smtp.SendMail(addr, a, from, to, doc.Bytes())
	return err
}
