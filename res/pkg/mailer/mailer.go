package mailer

const Mailer = `package mailer
import (
	"fmt"
	"net/smtp"
	"strings"

	"{{ .ModuleName}}/pkg/config"
	"{{ .ModuleName}}/pkg/logger"
)

// MailDetails struc
type MailDetails struct {
	To      []string
	Body    string
	Subject string
}

// SendMail function
func SendMail(mailDetails MailDetails) error {

	// Set up Init config
	smpt := config.SMTPCfg()
	addr := fmt.Sprintf("%s:%s", smpt.Host, smpt.Port)

	// set the sender, recipient, subject and body
	from := fmt.Sprintf("%s <%s>", smpt.From, smpt.From) // Add sender name
	msg := strings.Join([]string{
		"From: " + from,
		"To: " + strings.Join(mailDetails.To, ","),
		"MIME-Version: 1.0",
		"Content-Type: text/html; charset=UTF-8",
		"Subject: " + mailDetails.Subject,
		"",
		mailDetails.Body,
	}, "\r\n")

	// Authentication
	auth := smtp.PlainAuth("", smpt.User, smpt.Password, smpt.Host)

	// send the email
	err := smtp.SendMail(addr, auth, smpt.From, mailDetails.To, []byte(msg))
	if err != nil {
		logger.Error("error while sending mail", err)
		return err
	}

	logger.Info("mail sent successfully")
	return nil
}
`
