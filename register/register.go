package register

import (
	"encoding/json"
	"fmt"
	"net/http"

	entity "api/entity"
	"api/smtps"

	"golang.org/x/crypto/bcrypt"
)

func Create(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	to := r.FormValue("to")
	// message := r.FormValue("message")
	var user entity.User
	// var users = []entity.User{}
	json.NewDecoder(r.Body).Decode(&user)
	// here we can give the user Password in byte form in password variable
	password := []byte(string(user.Password))
	// here we can create hashed Passwword by bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		panic(err)
	}
	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	fmt.Println(err) // nil means it is a match

	// here we can convert hashed password in string form and store in user Password
	user.Password = string(hashedPassword)
	// database.Database.Find(&users)
	// fmt.Println(users)

	// here we can create the data on the database but the password will be saved in hashed Password
	// database.Database.Create(&user
	fmt.Println(to)
	smtps.Smtp(to, "this is my last message")
}

// func Smtp(to string, message string) {
// 	from := "aryanpal692@gmail.com"
// 	password := "uhmuvyuxufshuurw"

// 	// to := []string{
// 	// 	"mohan@mailinator.com",
// 	// }
// 	smtpHost := "smtp.gmail.com"
// 	smtpPort := "25"

// 	// message := []byte("This is a new test email message.")

// 	auth := smtp.PlainAuth("", from, password, smtpHost)

// 	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("Email Sent Successfully!")

// }
