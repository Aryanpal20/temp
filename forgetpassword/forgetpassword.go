package forgetpassword

import (
	"api/database"
	"api/entity"
	sms "api/messages"
	smtp "api/smtps"
	"net/http"
)

func Forgetpassword(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	var otp entity.Otp
	email := r.FormValue("email")
	database.Database.Where("email = ?", email).Find(&user)
	database.Database.Where("userid = ?", user.ID).Find(&otp)
	// here we can check user's email with user table database
	if email == user.Email {
		smtp.Smtp(user.Email, "Subject: ForgetPassword \r\n\r\n"+"Hi, "+user.Username+" \n OTP for change your password"+
			"\n OTP : "+otp.Otp)
		sms.SMS(user.Phone, "Hello, "+user.Username+" Enter the otp and change your password "+"\n OTP : "+otp.Otp)

	}
	// phone := r.FormValue("phone")
	// database.Database.Where("phone = ?", phone).Find(&user)
	// database.Database.Where("userid = ?", user.ID).Find(&otp)
	// if phone == user.Phone {
	// 	sms.SMS(user.Phone, "Hello, "+user.Username+" Enter the otp and change your password "+"\n OTP : "+otp.Otp)
	// }

}
