package helpers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"personal-growth/configs"
)

type RegistrationEmailData struct {
	Name     string
	AppName  string
	LoginURL string
	Otp      string
}

func RenderEmailTemplate(filename string, data RegistrationEmailData) (string, error) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		return "", err
	}
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return "", err
	}
	return body.String(), nil
}

func SendEmail(to string, subject string, body string) error {
	config, _ := configs.LoadConfig(".")
	// SMTP server config.
	smtpHost := "smtp.gmail.com" // Example: smtp.gmail.com
	smtpPort := "587"            // Usually 587 for TLS

	// Compose the message.
	message := []byte(fmt.Sprintf("From: KaKa Project <%s>\r\n", config.EmailAddress) +
		fmt.Sprintf("To: %s\r\n", to) +
		fmt.Sprintf("Subject: %s\r\n", subject) +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"utf-8\"\r\n" +
		"\r\n" +
		fmt.Sprintf("%s\r\n", body))

	// Auth.
	auth := smtp.PlainAuth("", config.EmailAddress, config.EmailPassword, smtpHost)

	// Send email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, config.EmailAddress, []string{to}, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return err
	}

	fmt.Println("Email is sent successfully!")

	return nil
}
