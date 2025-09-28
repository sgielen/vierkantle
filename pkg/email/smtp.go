package email

import (
	"flag"
	"fmt"
	"log/slog"
	"net/smtp"
)

var (
	smtp_host = flag.String("smtp_host", "", "SMTP host to relay e-mails through")
	smtp_port = flag.Uint("smtp_port", 25, "SMTP port")
	smtp_user = flag.String("smtp_user", "", "SMTP username for authentication (if given)")
	smtp_pass = flag.String("smtp_pass", "", "SMTP password for authentication")
	smtp_iden = flag.String("smtp_iden", "", "SMTP identity for authentication (usually empty)")

	smtp_from = flag.String("smtp_from", "app@vierkantle.nl", "Sender for e-mails")
)

var ErrNoSettings = fmt.Errorf("cannot send e-mail: missing settings")

func Init() {
	if err := CheckFlags(); err != nil {
		slog.Error("failed to initialize email", slog.Any("err", err))
	}
}

func CheckFlags() error {
	if *smtp_host == "" {
		return ErrNoSettings
	}
	return nil
}

func SendMail(to []string, msg []byte) error {
	if *smtp_host == "" {
		return ErrNoSettings
	}
	var auth smtp.Auth
	if *smtp_user != "" {
		auth = smtp.PlainAuth(*smtp_iden, *smtp_user, *smtp_pass, *smtp_host)
	}
	return smtp.SendMail(fmt.Sprintf("%s:%d", *smtp_host, *smtp_port), auth, *smtp_from, to, msg)
}
