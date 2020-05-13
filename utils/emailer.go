package utils

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"os"
)

// SendMail trough TLS with setuped mail from ENV
func SendMail(to string, body []byte) error {

	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		os.Getenv("MAIL_USERNAME"),
		os.Getenv("MAIL_PASSWORD"),
		os.Getenv("MAIL_HOST"),
	)

	host := os.Getenv("MAIL_HOST")
	addr := fmt.Sprintf("%s:%s", host, os.Getenv("MAIL_PORT"))
	from := os.Getenv("MAIL_USERNAME")

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = to
	headers["Subject"] = "test"
	headers["Content-Type"] = "text/html; charset=\"UTF-8\";"
	headers["MIME-version"] = "1.0"
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + string(body)

	tlsconfig := &tls.Config{ServerName: host}

	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return err
	}
	defer conn.Close()
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	defer c.Close()
	if err = c.Hello("localhost"); err != nil {
		return err
	}
	if err = c.Auth(auth); err != nil {
		return err
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	if err = c.Rcpt(to); err != nil {
		return err
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}
