package otp

import (
	"api/database"
	"api/entity"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func OTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	var otp entity.Otp
	var otps []entity.Otp
	json.NewDecoder(r.Body).Decode(&otp)
	database.Database.Where("email = ?", otp.Email).Find(&user)
	database.Database.Where("email = ?", otp.Email).Find(&otps)
	fmt.Println(len(otps))
	fmt.Println(user)
	var max = 9999
	var min = 1000
	v := rand.Intn(max-min) + min
	otp.Expire_time = time.Now().Local().Add(time.Minute * time.Duration(1))
	otp.Otp = strconv.Itoa(v)
	otp.Userid = user.ID
	if len(otps) == 1 {
		if err := database.Database.Model(&otp).Where("email = ?", user.Email).Update("otp", otp.Otp).Error; err != nil {
			fmt.Printf("update err != nil; %v\n", err)
		}
		if err := database.Database.Model(&otp).Where("email = ?", user.Email).Update("expire_time", otp.Expire_time).Error; err != nil {
			fmt.Printf("update err != nil; %v\n", err)
		}
		json.NewEncoder(w).Encode(otp)
	} else {
		database.Database.Create(&otp)
		json.NewEncoder(w).Encode(otp)
	}

}
