package services

import (
	"fmt"
	"net/smtp"
	"os"
)

// EmailSender define la interfaz para enviar correos electrónicos.
type EmailSender interface {
	SendEmail(to string, subject string, body string) error
}

// SMTPEmailSender es una implementación de EmailSender que usa SMTP.
type SMTPEmailSender struct {
	SMTPHost       string
	SMTPPort       string
	SenderEmail    string
	SenderPassword string
}

// NewSMTPEmailSender crea una nueva instancia de SMTPEmailSender.
func NewSMTPEmailSender() *SMTPEmailSender {
	return &SMTPEmailSender{
		SMTPHost:       os.Getenv("SMTP_HOST"),
		SMTPPort:       os.Getenv("SMTP_PORT"),
		SenderEmail:    os.Getenv("SENDER_EMAIL"),
		SenderPassword: os.Getenv("SENDER_PASSWORD"),
	}
}

// SendEmail envía un correo electrónico usando SMTP.
func (s *SMTPEmailSender) SendEmail(to string, subject string, body string) error {
	auth := smtp.PlainAuth("", s.SenderEmail, s.SenderPassword, s.SMTPHost)
	err := smtp.SendMail(s.SMTPHost+":"+s.SMTPPort, auth, s.SenderEmail, []string{to}, []byte("Subject: "+subject+"\r\n\r\n"+body))
	if err != nil {
		fmt.Println("Error enviando correo:", err)
		return err
	}
	fmt.Println("Correo enviado con exito.")
	return nil
}
