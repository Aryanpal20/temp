package smtps

import (
	"fmt"
	"net/smtp"
)

func Smtp(to string, message string) {
	from := "aryanpal692@gmail.com"
	password := "uhmuvyuxufshuurw"

	// to := []string{
	// 	"mohan@mailinator.com",
	// }
	smtpHost := "smtp.gmail.com"
	smtpPort := "25"

	// message = []byte("From: aryanpal692@gmail.com\r\n" +
	// 	"To: mohan@mailinator.com\r\n" +
	// 	"Subject: Test mail\r\n\r\n" +
	// 	"Email body\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")

}
