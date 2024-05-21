package service

import (
	"fmt"
	"net/smtp"

	"github.com/handarudwiki/golang-ewalet/config"
	"github.com/handarudwiki/golang-ewalet/domain"
)

type emailService struct {
	cnf *config.Config
}

func NewEmail(cnf *config.Config) domain.EmailService {
	return &emailService{
		cnf: cnf,
	}
}

func (e emailService) Send(to, subject, body string) error {
	auth := smtp.PlainAuth("", e.cnf.Email.Email, e.cnf.Email.Password, e.cnf.Email.Host)
	msg := []byte("From : handaru <" + e.cnf.Email.Email + ">\n" +
		"To : " + to + "\n" +
		"Subject : " + subject +
		"\nBody : " + body,
	)

	fmt.Println(to)
	fmt.Println(subject)
	fmt.Println(body)

	return smtp.SendMail(e.cnf.Email.Host+":"+e.cnf.Email.Port, auth, e.cnf.Email.Email, []string{to}, msg)
}
