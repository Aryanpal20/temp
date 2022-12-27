package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func Salary(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header["Token"][0], " ")[1]
	a := is_manager(token)
	var email string
	if a == "manager" {
		var user User
		var z int = 0
		var e int
		var task Task
		var tasks = []Task{}
		if user.Email == email {
			json.NewDecoder(r.Body).Decode(&user)
			Database.Where("email = ?", user.Email).Find(&user)
			if task.Assign == email {
				Database.Where("assign = ?", user.Email).Find(&tasks)
				// json.NewEncoder(w).Encode(tasks)
				for _, k := range tasks {
					Database.Where("assign = ?", user.Email).Find(&task)
					if k.Status == 1 {
						c := user.Hourly_Rate * k.Working_Hours
						if k.Estimate_time_work >= k.Work_Done_time {
							d := user.Hourly_Rate * k.Working_Hours
							e = d * 5 / 100
							fmt.Println("the bonous is : ", e)
						}
						z = z + c + e
					}

				}

			}
		}
		json.NewEncoder(w).Encode(z)
	} else {
		b := "access denied !!!"
		json.NewEncoder(w).Encode(b)
	}
}
