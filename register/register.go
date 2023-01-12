package register

import (
	"encoding/json"
	"fmt"
	"net/http"

	"api/database"
	entity "api/entity"

	sms "api/messages"
	"api/smtps"

	"golang.org/x/crypto/bcrypt"
)

func Create(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	var users = []entity.User{}
	// here we want to decode the user
	json.NewDecoder(r.Body).Decode(&user)
	// here we want to fetch the data from database as user input
	database.Database.Where("email = ?", user.Email).Find(&users)
	// here we can give the condition.
	if len(users) == 0 {
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

		// // here we can create the data on the database but the password will be saved in hashed Password
		database.Database.Create(&user)

		// fmt.Println(password)
		// here we want to send email message on user email.
		smtps.Smtp(user.Email, "Subject: Registration \r\n\r\n"+"Hi, "+user.Username+" \nYour Account has been created successfully"+
			"\n Email : "+user.Email+"\n Password : "+string(password))
		// fmt.Println(user.Phone)
		// here we want to send a sms message on user phone number
		sms.SMS(user.Phone, "Hello, "+user.Username+" Your Account has been created successfully")
	} else {
		b := "This Email already exist"
		json.NewEncoder(w).Encode(b)
	}

}

// ----------Hard Code for SMS-----------------

// func Smtp() {
// 	from := "aryanpal692@gmail.com"
// 	password := "uhmuvyuxufshuurw"

// to := []string{
// 	"mohan@mailinator.com",
// }
// 	smtpHost := "smtp.gmail.com"
// 	smtpPort := "25"

// message := []byte("This is a new test email message.")

// 	auth := smtp.PlainAuth("", from, password, smtpHost)

// 	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("Email Sent Successfully!")

// }
