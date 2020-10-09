package email

import (
	"base-site-api/internal/app/config"
	"crypto/tls"
	"fmt"
	"net/smtp"
)

type Emailer interface {
	SendMail(to string, body []byte) error
}

var client Emailer

func Initialize(c *config.EmailConfiguration) {
	client = &SmtpClient{c}
}

func Instance() Emailer {
	return client
}

type SmtpClient struct {
	c *config.EmailConfiguration
}

type clientTLS struct {
	smtpClient *smtp.Client
	tlsConn    *tls.Conn
}

// SendMail trough TLS with setuped mail from ENV
func (s SmtpClient) SendMail(to string, body []byte) error {
	host := s.c.SmtpHost
	addr := fmt.Sprintf("%s:%s", host, s.c.Port)
	from := s.c.Username

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

	client, err := s.setupClient(host, addr)

	if err != nil {
		return err
	}
	c := client.smtpClient

	defer client.tlsConn.Close()
	defer c.Close()

	err = setupFromTo(c, from, to)

	if err != nil {
		return err
	}

	err = writeData(c, message)

	if err != nil {
		return err
	}

	return c.Quit()
}

func (s SmtpClient) setupClient(host string, addr string) (*clientTLS, error) {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		s.c.Username,
		s.c.Password,
		s.c.SmtpHost,
	)

	tlsconfig := &tls.Config{ServerName: host}

	conn, err := tls.Dial("tcp", addr, tlsconfig)
	if err != nil {
		return nil, err
	}

	// defer conn.Close()

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return nil, err
	}

	if err = c.Auth(auth); err != nil {
		return nil, err
	}

	return &clientTLS{
		c,
		conn,
	}, nil
}

func writeData(c *smtp.Client, message string) error {
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

	return nil
}

func setupFromTo(c *smtp.Client, from string, to string) error {
	if err := c.Mail(from); err != nil {
		return err
	}

	if err := c.Rcpt(to); err != nil {
		return err
	}

	return nil
}
