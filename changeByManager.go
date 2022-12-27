package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func ChangeByManager(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header["Token"][0], " ")[1]
	a := is_manager(token)
	if a == "manager" {
		var tasks Task
		Database.First(&tasks, mux.Vars(r)["id"])
		json.NewDecoder(r.Body).Decode(&tasks)
		if err := Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("assign", tasks.Assign).Error; err != nil {
			fmt.Printf("update err != nil; %v\n", err)
		}
		if err := Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("estimate_time_work", tasks.Estimate_time_work).Error; err != nil {
			fmt.Printf("update err != nil; %v\n", err)
		}

		// Database.Create(&task)
		json.NewEncoder(w).Encode(tasks)
	} else {
		b := "access denied !!!"
		json.NewEncoder(w).Encode(b)
	}

}
