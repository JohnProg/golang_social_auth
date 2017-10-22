package mailers

import (
	"crypto/tls"
	"html/template"
	"time"

	"bytes"
	"fmt"

	"log"

	"gopkg.in/gomail.v2"
	"github.com/golang_social_auth/settings"
)

var templateFunctions template.FuncMap

func init() {
	templateFunctions = template.FuncMap{
		"formatDate": formatDate,
	}
}

func Send(to, subject, body string) error {
	d := gomail.NewPlainDialer(
		settings.Config.SMTP.Host,
		settings.Config.SMTP.Port,
		settings.Config.SMTP.Username,
		settings.Config.SMTP.Password,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", settings.Config.SMTP.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return d.DialAndSend(m)
}

func RenderTemplate(tpl string, binding interface{}) (string, error) {
	var html bytes.Buffer

	tmpl, err := template.New("layout.tmpl").Funcs(templateFunctions).ParseFiles("templates/mails/layout.tmpl", fmt.Sprintf("templates/mails/%s.tmpl", tpl))
	if err != nil {
		log.Printf("MAILER: Rendering Error:  %e", err)
		return "", err
	}

	err = tmpl.Execute(&html, binding)
	if err != nil {
		log.Printf("MAILER: Rendering Error:  %e", err)
		return "", err
	}

	return html.String(), nil
}

func formatDate(timestamp int64) string {
	t := time.Unix(timestamp/1000, 0)
	return t.Format(time.UnixDate)
}
