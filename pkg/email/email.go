package email

import (
	"context"
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	from = mail.NewEmail("Vierkantle", "info@vierkantle.nl")
)

func GetSendgridApiKey() string {
	return os.Getenv("SENDGRID_API_KEY")
}

func SendLoginUrl(ctx context.Context, email string, url string) error {
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
	return sendMail(ctx, message)
}

func SendNewBoardInQueue(ctx context.Context, author string, name string) error {
	subject := "Nieuw bord in wachtrij"
	plainTextContent := `
Er staat een nieuw bord in de wachtrij van ` + author + `! Naam is ` + name + `.

Zie: https://vierkantle.nl/generator
`

	message := mail.NewSingleEmail(from, subject, from, plainTextContent, "")
	return sendMail(ctx, message)
}

func sendMail(ctx context.Context, message *mail.SGMailV3) error {
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
