// here we can use foreign key
package taskcreate1

import (
	"api/database"
	"api/entity"
	email "api/fetchemail"
	role "api/fetchrole"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func TaskCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	token := strings.Split(r.Header["Token"][0], " ")[1]
	// here we can store the value of role in a variable
	a := role.Is_manager(token)
	// here we can store the value of email in c variable
	c := email.Task_creator(token)
	rep := c
	if a == "manager" {
		var task entity.Task
		json.NewDecoder(r.Body).Decode(&task)
		database.Database.Where("id = ?", task.UserId).Find(&user)
		fmt.Println(user)
		task.Assign = user.Email
		// here we can assign the value of email which is present in token to task reportor
		task.Reportor = rep
		// here it will take current time and stored in task table field in Created_At
		task.Created_At = time.Now().String()
		// database.Database.Create(&task)
		json.NewEncoder(w).Encode(task)
	} else {
		b := "access denied !!!"
		json.NewEncoder(w).Encode(b)
	}

}
