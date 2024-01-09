package utils

import (
	"github.com/delta/FestAPI/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(toName string, toEmail string, subject string, plainTextContent string, htmlContent string) error {

	from := mail.NewEmail(config.SenderName, config.SenderEmail)
	to := mail.NewEmail(toName, toEmail)

	print(config.SenderName, config.SenderEmail)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(config.SendgridAPIKey)

	_, err := client.Send(message)

	if err != nil {
		return err
	}
	return nil
}
