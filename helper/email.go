package helper

import (
	"bytes"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
)

func SendEmail(subject string, email string, message string) {
	const CONFIG_SMTP_HOST = "mail.masuk.email"
	const CONFIG_SMTP_PORT = 465
	const CONFIG_SENDER_NAME = "Service Myfin <service@myfin.id>"
	const CONFIG_AUTH_EMAIL = "service@myfin.id"
	const CONFIG_AUTH_PASSWORD = "Adminmyfin123"

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", message)

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)
	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

}

func EmailTemplate(writer bytes.Buffer, templateOTP string, otp string) string {
	t, err := template.New("OTP").Parse(templateOTP)
	PanicIfError(err)

	t.ExecuteTemplate(&writer, "OTP", map[string]interface{}{
		"Otp": otp,
	})
	return writer.String()
}
