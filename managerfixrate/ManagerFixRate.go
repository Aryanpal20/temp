package managerfixrate

import (
	database "api/database"
	entity "api/entity"
	role "api/fetchrole"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func ManagerFixRate(w http.ResponseWriter, r *http.Request) {
	// here we can split the token and decode the token
	token := strings.Split(r.Header["Token"][0], " ")[1]
	a := role.Is_manager(token)
	var email string
	// here we can compare with manager
	if a == "manager" {
		var user entity.User
		if user.Email == email {
			json.NewDecoder(r.Body).Decode(&user)
			// here we can update the hourly rate by manager
			if err := database.Database.Model(&user).Where("email = ?", user.Email).Update("hourly_rate", user.Hourly_Rate).Error; err != nil {
				fmt.Printf("update err != nil; %v\n", err)
			}
		}
		database.Database.Save(user)
		json.NewEncoder(w).Encode(user)
	} else {
		b := "access denied !!!"
		json.NewEncoder(w).Encode(b)
	}
}
