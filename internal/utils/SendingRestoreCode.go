package utils

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// smtpServer data to smtp server
type smtpServer struct {
	host string
	port string
} // Address URI to smtp server

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

//StringWithCharset needs to generate Code
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

//String needs to generate code with given Lenght
func String(length int) string {
	return StringWithCharset(length, charset)
}

// SendRestoreCodeToUser - needs to allow restore password if user fargot it
func SendRestoreCodeToUser(mail, code string) {
	from := "forum.bot.alem@gmail.com"
	password := "8LYL6GnWpwqgs3k" // Receiver email address.
	to := []string{
		mail,
	} // smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	msg := "Here your code to restore your password: " + code   // Message.
	message := []byte(msg)                                      // Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host) // Sending email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		fmt.Println("Sending mail is broken srry")
		return
	}
}
