package changebyemployee

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	database "api/database"
	entity "api/entity"
	email "api/fetchemail"
	role "api/fetchrole"

	"github.com/gorilla/mux"
)

func ChangeByEmployee(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header["Token"][0], " ")[1]
	a := role.Is_manager(token)
	// here we can compare with employee
	if a == "employee" {
		var tasks entity.Task
		// here we can find the from database by id
		database.Database.First(&tasks, mux.Vars(r)["id"])
		c := email.Task_creator(token)
		rep := c
		if rep == tasks.Assign {
			// Database.First(&tasks, mux.Vars(r)["id"])
			json.NewDecoder(r.Body).Decode(&tasks)
			// here we can update the comment by employee
			if err := database.Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("comment", tasks.Comment).Error; err != nil {
				fmt.Printf("update err != nil; %v\n", err)
			}
			// here we can update the status by employee
			if err := database.Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("status", tasks.Status).Error; err != nil {
				fmt.Printf("update err != nil; %v\n", err)
			}
			// here we can update the working hour by employee
			if err := database.Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("working_hours", tasks.Working_Hours).Error; err != nil {
				fmt.Printf("update err != nil; %v\n", err)
			}
			// here we can update the work done time by employee
			if err := database.Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("work_done_time", tasks.Work_Done_time).Error; err != nil {
				fmt.Printf("update err != nil; %v\n", err)
			}

			// Database.Create(&task)
			json.NewEncoder(w).Encode(tasks)
		} else {
			d := "you can't change"
			json.NewEncoder(w).Encode(d)
		}

	} else {
		b := "access denied !!!"
		json.NewEncoder(w).Encode(b)
	}

}
