package main

import (
	"gopkg.in/gomail.v2"
)

func SendEmail(to string, header string, contents string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.Smtp.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", header)
	m.SetBody("text/plain", contents)

	d := gomail.Dialer{Host: config.Smtp.Server, Port: config.Smtp.Port}
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
