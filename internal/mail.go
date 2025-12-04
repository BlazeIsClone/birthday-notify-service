package internal

import (
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendMail(toName string, toEmail string, subject string, message string) {
	from := mail.NewEmail(os.Getenv("SENDGRID_FROM_NAME"), os.Getenv("SENDGRID_FROM_EMAIL"))
	to := mail.NewEmail(toName, toEmail)
	email := mail.NewSingleEmail(from, subject, to, message, message)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	response, err := client.Send(email)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		return
	}

	log.Printf("Mail Sent with Status Code: %d", response.StatusCode)
}
