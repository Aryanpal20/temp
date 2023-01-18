package salary

import (
	database "api/database"
	entity "api/entity"
	role "api/fetchrole"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func Salary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := strings.Split(r.Header["Token"][0], " ")[1]
	a := role.Is_manager(token)
	var email string
	// here we can compare with manager as well as admin
	if a == "manager" || a == "admin" {
		var user entity.User
		var z int = 0
		var e int
		var task entity.Task
		// here we can use tasks variable to store variable
		var tasks = []entity.Task{}
		if user.Email == email {
			json.NewDecoder(r.Body).Decode(&user)
			database.Database.Where("email = ?", user.Email).Find(&user)
			if task.Assign == email {
				database.Database.Where("assign = ?", user.Email).Find(&tasks)
				// json.NewEncoder(w).Encode(tasks)
				// here we can use for loop for range in tasks for get all same email details
				for _, k := range tasks {
					database.Database.Where("assign = ?", user.Email).Find(&task)
					// here we can compare status is equal to completed
					if k.Status == 1 || k.Status == 2 || k.Status == 0 {
						// here we get a salary of completed tasks
						c := user.Hourly_Rate * k.Working_Hours
						z = z + c
					}
					// here we can check for bonus if work is done before estimate time given by manager
					if k.Estimate_time_work >= k.Work_Done_time {
						d := user.Hourly_Rate * k.Working_Hours
						// here we can give the bonus of 5%.
						e = d * 5 / 100
						fmt.Println("the bonous is : ", e)
						// here swe can add the bonus in the employee salary
						z = z + e
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
