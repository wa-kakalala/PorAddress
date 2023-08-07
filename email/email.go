package email

import (
	"gopkg.in/gomail.v2"
)

type EmailInfo struct {
	To      string
	From    string
	Subject string
	Host    string
	Port    int
	User    string
	Passwd  string
}

// https://github.com/go-gomail/gomail
func (ei *EmailInfo) SendEmail(content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", ei.From)
	m.SetHeader("To", ei.To) //主送
	m.SetHeader("Subject", ei.Subject)
	m.SetBody("text/plain", content)
	d := gomail.NewDialer(ei.Host, ei.Port, ei.User, ei.Passwd)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
