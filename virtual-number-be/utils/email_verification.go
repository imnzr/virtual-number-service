package utils

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"os"
	"time"
)

func SendEmail(to, subject, body string) error {
	from := os.Getenv("SMTP_SENDER")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	addr := smtpHost + ":" + smtpPort

	// Dial koneksi TCP
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		return fmt.Errorf("gagal koneksi SMTP: %w", err)
	}

	// Buat SMTP client
	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return fmt.Errorf("gagal buat SMTP client: %w", err)
	}
	defer client.Close()

	// Jalankan STARTTLS
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true, // opsional: abaikan sertifikat invalid
		ServerName:         smtpHost,
	}

	if err = client.StartTLS(tlsconfig); err != nil {
		return fmt.Errorf("STARTTLS gagal: %w", err)
	}

	// Auth setelah STARTTLS
	auth := smtp.PlainAuth("", from, password, smtpHost)
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("auth gagal: %w", err)
	}

	// Kirim email
	if err = client.Mail(from); err != nil {
		return err
	}
	if err = client.Rcpt(to); err != nil {
		return err
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	message := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, to, subject, body))

	_, err = w.Write(message)
	if err != nil {
		return err
	}
	w.Close()

	return client.Quit()
}
