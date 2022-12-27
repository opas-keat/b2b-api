package utils

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"log"
)

func SendMail2() bool {
	m := gomail.NewMessage()
	// Set E-Mail sender
	m.SetHeader("From", "ppsuperwheel.b2b@gmail.com")
	// Set E-Mail receivers
	m.SetHeader("To", "opas.keat@gmail.com")

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject2")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "This is Gomail test body2")

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.mailtrap.io", 587, "cacae0b9ea4212", "e051356a50db51")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Could not send email: ", err)
		return false
	}

	return true
}
