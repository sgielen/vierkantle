package email

import (
	"context"
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func GetSendgridApiKey() string {
	return os.Getenv("SENDGRID_API_KEY")
}

func SendLoginUrl(ctx context.Context, email string, url string) error {
	from := mail.NewEmail("Vierkantle", "info@vierkantle.nl")
	subject := "Je Vierkantle-login"
	to := mail.NewEmail(email, email)
	plainTextContent := `
Iemand, waarschijnlijk jij, heeft geprobeerd in te loggen op Vierkantle met dit e-mailadres.
Als jij dat niet was, negeer dan vooral dit mailtje, dan gebeurt er niets.

Was jij het wel, gebruik dan de volgende URL om in te loggen:
	` + url + `

Veel plezier met Vierkantlen!
`
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, "")
	client := sendgrid.NewSendClient(GetSendgridApiKey())
	response, err := client.SendWithContext(ctx, message)
	if err != nil {
		return fmt.Errorf("sending e-mail failed because of %w", err)
	}
	if response.StatusCode >= 300 {
		return fmt.Errorf("sending e-mail failed because of HTTP status %d %q", response.StatusCode, response.Body)
	}
	return nil
}
