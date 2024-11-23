package mail

import (
	"fmt"
	"net/smtp"
	"strings"

	"app/internal/config"
	"app/internal/log"
)

func SendEmail(to string, subject string, body string) error {
	host := config.C.Mail.Host
	port := config.C.Mail.Port
	smtpServer := fmt.Sprintf("%s:%d", host, port)
	username := config.C.Mail.Username
	password := config.C.Mail.Password

	auth := smtp.PlainAuth("", username, password, host)

	message := strings.Builder{}
	message.WriteString(fmt.Sprintf("From: %s\r\n", username))
	message.WriteString(fmt.Sprintf("To: %s\r\n", to))
	message.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))
	message.WriteString(
		"MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";",
	)
	message.WriteString("\r\n")
	message.WriteString(body)

	log.S.Debug("Sending confirmation email", log.L())
	err := smtp.SendMail(
		smtpServer,
		auth,
		username,
		[]string{to},
		[]byte(message.String()),
	)
	if err != nil {
		log.S.Info("Failed to send email", log.L().Error(err))
		return err
	}

	log.S.Debug("Successfully sent confirmation email", log.L())
	return nil
}
