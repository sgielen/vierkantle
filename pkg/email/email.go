package email

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"time"
)

var (
	from = "Vierkantle <info@vierkantle.nl>"
)

func SendLoginUrl(ctx context.Context, email string, url string) error {
	subject := "Je Vierkantle-login"
	plainTextContent := `
Iemand, waarschijnlijk jij, heeft geprobeerd in te loggen op Vierkantle met dit e-mailadres.
Als jij dat niet was, negeer dan vooral dit mailtje, dan gebeurt er niets.

Was jij het wel, gebruik dan de volgende URL om in te loggen:
	` + url + `

Veel plezier met Vierkantlen!
`
	return sendMail(ctx, email, subject, plainTextContent)
}

func SendNewBoardInQueue(ctx context.Context, author string, name string) error {
	subject := "Nieuw bord in wachtrij"
	plainTextContent := `
Er staat een nieuw bord in de wachtrij van ` + author + `! Naam is ` + name + `.

Zie: https://vierkantle.nl/generator
`

	return sendMail(ctx, from, subject, plainTextContent)
}

func sendMail(ctx context.Context, to string, subject string, content string) error {
	var buffer bytes.Buffer

	if err := write(&buffer,
		fmt.Sprintf("Subject: %s\r\n", subject),
		fmt.Sprintf("From: %s\r\n", from),
		fmt.Sprintf("To: %s\r\n", to),
		fmt.Sprintf("Date: %s\r\n", time.Now().Format(time.RFC1123Z)),
		"\r\n\r\n",
		content); err != nil {
		return err
	}

	err := SendMail([]string{to}, buffer.Bytes())
	if err != nil {
		slog.ErrorContext(ctx, "An e-mail failed to be sent",
			slog.String("email", to),
			slog.String("subject", subject),
			slog.Any("err", err))
		return err
	} else {
		slog.InfoContext(ctx, "An e-mail was sent",
			slog.String("email", to),
			slog.String("subject", subject))
		return nil
	}
}

func write(w io.Writer, parts ...string) error {
	for _, p := range parts {
		_, err := w.Write([]byte(p))
		if err != nil {
			return err
		}
	}
	return nil
}
