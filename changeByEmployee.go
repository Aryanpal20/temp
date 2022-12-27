package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func ChangeByEmployee(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header["Token"][0], " ")[1]
	a := is_manager(token)
	if a == "employee" {
		var tasks Task
		Database.First(&tasks, mux.Vars(r)["id"])
		c := task_creator(token)
		rep := c
		if rep == tasks.Assign {
			// Database.First(&tasks, mux.Vars(r)["id"])
			json.NewDecoder(r.Body).Decode(&tasks)
			if err := Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("comment", tasks.Comment).Error; err != nil {
				fmt.Printf("update err != nil; %v\n", err)
			}
			if err := Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("status", tasks.Status).Error; err != nil {
				fmt.Printf("update err != nil; %v\n", err)
			}
			if err := Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("working_hours", tasks.Working_Hours).Error; err != nil {
				fmt.Printf("update err != nil; %v\n", err)
			}
			if err := Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("work_done_time", tasks.Work_Done_time).Error; err != nil {
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
