package app

import (
	"async-1/mail-services/config"
	"bytes"
	"path"
	"text/template"

	"gopkg.in/gomail.v2"
)

type repository struct {
	email config.EmailConfig
}

func NewRepository(email config.EmailConfig) repository {
	return repository{
		email: email,
	}
}

func (e *repository) SendEmail(email Email) (err error) {
	filepath := path.Join("web", "email.html")
	tmpl := template.Must(template.ParseFiles(filepath))

	bufer := new(bytes.Buffer)
	err = tmpl.Execute(bufer, email)
	if err != nil {
		return
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", e.email.SenderName)
	mailer.SetHeader("To", email.To...)
	mailer.SetHeader("Subject", email.Subject)
	mailer.SetBody(email.Type, email.Message)

	dialer := gomail.NewDialer(
		e.email.SmtpHost,
		e.email.SmtpPort,
		e.email.AuthEmail,
		e.email.AuthPassword,
	)

	return dialer.DialAndSend(mailer)
}
