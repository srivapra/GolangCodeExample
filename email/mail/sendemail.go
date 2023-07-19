package mail

import (
	"fmt"
	"net/smtp"
)

// SendEmail() function to send the email
func SendEmail() {

	// Sender email and password which is needed to send the email.
	from := "prashant.srivastava7744@gmail.com"
	password := "krzkoafadybwsocd"

	// Receiver email address, We can add multiple email address also.
	to := []string{
		"javastudypoint98@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// The email that I am going to send along with subject line and body.
	subject := "Subject : Code to send email using Golang\n"
	body := "Hi, \n Please find the attached code to send am email using Golang.\n Thanks & Regards\n Prashant Srivastava"
	message := []byte(subject + body)

	// Creating authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending the email using SendMail function.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	// Handling the error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
