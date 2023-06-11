package utils

import (
	"net/mail"
	"net/smtp"
)

func SendEmail(
	from string,
	to []string,
	cc []string,
	bcc []string,
	subject string,
	body string,
	attachment string,
) error {
	host := "smtp.gmail.com"
	port := "587"
	subject = "Subject: " + subject + "\n"
	msg := []byte(subject + body)
	addr := host + ":" + port
	auth := smtp.PlainAuth("", from, "password", host)
	return smtp.SendMail(addr, auth, from, to, msg)
}

func ValidEmail(email string) (string, bool) {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}
