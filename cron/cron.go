package cron

import (
	"api/database"
	"api/entity"
	"encoding/json"
	"fmt"
	"net/http"
)

func Data(w http.ResponseWriter, r *http.Request) {

	var email = []string{"sonu@gmail.com", "rohan@gmail.com"}
	fmt.Println(email)
	for _, i := range email {
		if i == "sonu@gmail.com" || i == "rohan@gmail.com" {
			var tasks = []entity.Task{}
			database.Database.Where("assign = ?", i).Find(&tasks)
			json.NewEncoder(w).Encode(tasks)
			fmt.Println(tasks)
		}
	}
}
