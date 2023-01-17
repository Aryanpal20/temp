package varifyotp

import (
	"api/database"
	"api/entity"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func VerifyOTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	email := r.FormValue("email")
	otp1 := r.FormValue("otp")
	var otps entity.Otp
	var user entity.User
	database.Database.Where("otp = ?", otp1).Find(&otps)
	database.Database.Where("email = ? ", otps.Email).Find(&user)
	if email == user.Email {
		if otp1 == otps.Otp {
			if otps.Expire_time.After(time.Now()) {
				password := r.FormValue("password")
				confirmpassword := r.FormValue("confirmpassword")
				if password == confirmpassword {
					password1 := []byte(string(password))
					hashedPassword, err := bcrypt.GenerateFromPassword(password1, 10)
					if err != nil {
						panic(err)
					}
					err = bcrypt.CompareHashAndPassword(hashedPassword, password1)
					fmt.Println(err) // nil means it is a match
					user.Password = string(hashedPassword)
					if err := database.Database.Model(&user).Where("id = ?", otps.Userid).Update("password", user.Password).Error; err != nil {
						fmt.Printf("update err != nil; %v\n", err)
					}
				} else {
					c := "password not matched"
					json.NewEncoder(w).Encode(c)
				}

			} else {
				d := "your otp time expired"
				json.NewEncoder(w).Encode(d)
			}
		}
	} else {
		b := "Invalid OTP"
		json.NewEncoder(w).Encode(b)
	}

}
