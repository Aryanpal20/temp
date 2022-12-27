package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func ManagerFixRate(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header["Token"][0], " ")[1]
	a := is_manager(token)
	var email string
	if a == "manager" {
		var user User
		if user.Email == email {
			json.NewDecoder(r.Body).Decode(&user)
			if err := Database.Model(&user).Where("email = ?", user.Email).Update("hourly_rate", user.Hourly_Rate).Error; err != nil {
				fmt.Printf("update err != nil; %v\n", err)
			}
		}
		Database.Save(user)
		json.NewEncoder(w).Encode(user)
	} else {
		b := "access denied !!!"
		json.NewEncoder(w).Encode(b)
	}
}
