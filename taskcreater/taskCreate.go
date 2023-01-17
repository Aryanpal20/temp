// here we can't use foreign key
package taskcreater

import (
	database "api/database"
	entity "api/entity"
	email "api/fetchemail"
	role "api/fetchrole"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func Task_Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// here we can give the token in header for decode using Bearer
	token := strings.Split(r.Header["Token"][0], " ")[1]
	// here we can store the value of role in a variable
	a := role.Is_manager(token)
	// here we can store the value of email in c variable
	c := email.Task_creator(token)
	rep := c
	if a == "manager" {
		var task entity.Task
		json.NewDecoder(r.Body).Decode(&task)
		// here we can assign the value of email which is present in token to task reportor
		task.Reportor = rep
		// here it will take current time and stored in task table field in Created_At
		task.Created_At = time.Now().String()
		database.Database.Create(&task)
		json.NewEncoder(w).Encode(task)
	} else {
		b := "access denied !!!"
		json.NewEncoder(w).Encode(b)
	}

}
