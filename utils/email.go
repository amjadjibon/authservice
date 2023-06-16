package utils

import (
	"net/mail"
	"net/smtp"
)

func SendEmail(
	to string,
	subject string,
	body string,
) error {
	from := ""
	password := ""
	host := ""
	port := ""
	subject = "Subject: " + subject + "\n"
	msg := []byte(subject + body)
	addr := host + ":" + port
	auth := smtp.PlainAuth("", from, password, host)
	return smtp.SendMail(addr, auth, from, []string{to}, msg)
}

func ValidEmail(email string) (string, bool) {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}
